package filesystem

import (
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type Node struct {
	Name     string
	IsDir    bool
	Children []*Node
}

func buildTree(path string) (*Node, error) {
	pathInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("error accessing root path %s: %v", path, err)
		return nil, err
	}

	rootNode := &Node{
		Name:  pathInfo.Name(),
		IsDir: pathInfo.IsDir(),
	}

	if pathInfo.IsDir() {
		err = buildTreeRecursive(path, rootNode)
		if err != nil {
			log.Printf("error building tree %s: %v", path, err)
			return rootNode, nil
		}
	}

	return rootNode, nil
}

func buildTreeRecursive(currentPath string, currentNode *Node) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.EACCES {
			log.Printf("permissions denied  %s. Ignoring...", currentPath)
			return nil
		}
		return err
	}

	for _, entry := range entries {
		name := entry.Name()

		if name == "." || name == ".." || name == "mnt" {
			continue
		}

		entryPath := filepath.Join(currentPath, name)
		entryInfo, err := os.Stat(entryPath)
		if err != nil {
			if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.EACCES {
				log.Printf("Permission denied %s. ignoring...", entryPath)
				continue
			}
			continue
		}

		node := &Node{
			Name:  name,
			IsDir: entryInfo.IsDir(),
		}

		if entryInfo.IsDir() {
			err = buildTreeRecursive(entryPath, node)
			if err != nil {
				log.Printf("Error processing directories %s: %v. ignoring childrens...", entryPath, err)
			}
		}

		currentNode.Children = append(currentNode.Children, node)
	}

	return nil
}

func ListAll(path string) (*Node, error) {

	tree, err := buildTree(path)
	if err != nil {
		log.Printf("error building Tree %s: %v", path, err)
		return nil, err
	}

	fakeStorage := NewFakeStorage()
	fakeStorage.SaveTree(tree)

	return tree, nil
}
