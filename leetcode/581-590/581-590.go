package _581_590

import "strings"

// leetcode problem No. 588

type FsTrie struct {
	isFile   bool // is file or directory
	content  strings.Builder
	children map[string]*FsTrie
}

func newTrie() *FsTrie {
	return &FsTrie{children: make(map[string]*FsTrie)}
}

func (t *FsTrie) insert(path string, isFile bool) *FsTrie {
	curNode := t
	paths := strings.Split(path, "/")
	for _, p := range paths[1:] {
		if _, ok := curNode.children[p]; !ok {
			curNode.children[p] = newTrie()
		}
		curNode = curNode.children[p]
	}
	curNode.isFile = isFile
	return curNode
}

func (t *FsTrie) search(path string) *FsTrie {
	if path == "/" {
		return t
	}
	paths := strings.Split(path, "/")
	curNode := t
	for _, p := range paths[1:] {
		if _, ok := curNode.children[p]; !ok {
			return nil
		}
		curNode = curNode.children[p]
	}
	return curNode
}

type FileSystem struct {
	root *FsTrie
}

func Constructor() FileSystem {
	return FileSystem{root: newTrie()}
}

func (fs *FileSystem) Ls(path string) []string {
	t := fs.root.search(path)
	if t == nil {
		return nil
	}
	if t.isFile {
		return []string{path}
	}
	result := make([]string, 0, len(t.children))
	for k, _ := range t.children {
		result = append(result, k)
	}
	return result
}

func (fs *FileSystem) Mkdir(path string) {
	fs.root.insert(path, false)
}

func (fs *FileSystem) AddContentToFile(filePath string, content string) {
	node := fs.root.insert(filePath, true)
	node.content.WriteString(content)
}

func (fs *FileSystem) ReadContentFromFile(filePath string) string {
	node := fs.root.search(filePath)
	if node == nil {
		return ""
	}
	return node.content.String()
}
