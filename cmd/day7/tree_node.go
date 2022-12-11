package main

import (
	"errors"
	"fmt"
)

type treeNode struct {
	Parent   *treeNode
	Name     string
	Files    map[string]*file
	Children map[string]*treeNode
}

func (n *treeNode) AddFile(f *file) error {
	if n.Files[f.Name] != nil {
		return errors.New(fmt.Sprintf("File %s exists in dir %s", f.Name, n.Name))
	}
	n.Files[f.Name] = f

	return nil
}

func (n *treeNode) FileExists(name string) bool {
	return n.Files[name] != nil
}

func (n *treeNode) Mkdir(name string) (*treeNode, error) {
	if n.Children[name] != nil {
		return nil, errors.New(fmt.Sprintf("Directory %s exists in dir %s", name, n.Name))
	}
	node := &treeNode{
		Parent:   n,
		Name:     name,
		Files:    make(map[string]*file),
		Children: make(map[string]*treeNode),
	}
	n.Children[name] = node

	return node, nil
}

func (n *treeNode) Size() int {
	sz := 0
	for _, file := range n.Files {
		sz += file.Size
	}
	for _, child := range n.Children {
		sz += child.Size()
	}
	return sz
}
