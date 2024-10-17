package filesystem

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"sync"
)

type FakeStorage interface {
	ResetTree()
	SaveTree(tree *Node)
	DisplayTree(path string) []string
	SearchFile(ctx context.Context, file string) ([]string, error)
}

type fakeStorage struct {
	tree *Node
	mu   sync.RWMutex
}

var (
	instance FakeStorage
	once     sync.Once
)

func NewFakeStorage() FakeStorage {
	once.Do(func() {
		instance = &fakeStorage{}
	})
	return instance
}

func (f *fakeStorage) ResetTree() {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.tree = nil
	log.Println("FakeStorage: Tree has been reset.")
}

func (f *fakeStorage) SaveTree(tree *Node) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.tree = tree
	log.Println("FakeStorage: Tree has been saved.")
}

func (f *fakeStorage) DisplayTree(path string) []string {
	if path == "" {
		path = "/"
	}

	directories := f.findNodeNames(path)
	return directories
}

func (f *fakeStorage) findNodeNames(path string) []string {
	f.mu.RLock()
	defer f.mu.RUnlock()

	if f.tree == nil {
		fmt.Println("tree not initialized")
		return nil
	}

	node := f.findNodeByPath(path)
	if node == nil {
		fmt.Printf("find no nodes: %s\n", path)
		return nil
	}

	var directories []string
	for _, child := range node.Children {
		directories = append(directories, child.Name)
	}

	log.Printf("FakeStorage: find  %dchildren for the path  %s.", len(directories), path)
	return directories
}

func (f *fakeStorage) findNodeByPath(path string) *Node {
	path = filepath.Clean(path)
	pathParts := strings.Split(path, string(filepath.Separator))

	if path == "/" {
		return f.tree
	}

	currentNode := f.tree
	startIndex := 0

	if len(pathParts) > 0 && pathParts[0] == currentNode.Name {
		startIndex = 1
	}

	for _, part := range pathParts[startIndex:] {
		if part == "" || part == "." {
			continue
		}

		found := false
		for _, child := range currentNode.Children {
			if child.Name == part {
				currentNode = child
				found = true
				break
			}
		}

		if !found {
			return nil
		}
	}

	return currentNode
}

func (f *fakeStorage) SearchFile(ctx context.Context, file string) ([]string, error) {
	var wg sync.WaitGroup
	var foundFiles []string
	var goFuncCounter int

	var searchNode func(nodes []*Node, path string)
	searchNode = func(nodes []*Node, path string) {
		defer wg.Done()

		for _, node := range nodes {
			select {
			case <-ctx.Done():
				return
			default:
				currentPath := path + "/" + node.Name
				if node.Name == file {
					foundFiles = append(foundFiles, currentPath)
				}
				if node.IsDir {
					wg.Add(1)
					go searchNode(node.Children, currentPath)
				}

				goFuncCounter++
			}
		}
	}

	wg.Add(1)
	go searchNode(f.tree.Children, "")

	wg.Wait()

	if len(foundFiles) == 0 {
		return nil, fmt.Errorf("file does not exist")
	}
	log.Printf("Go func was called %d", goFuncCounter)

	return foundFiles, nil
}
