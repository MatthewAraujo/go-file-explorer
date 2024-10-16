package filesystem

import (
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
