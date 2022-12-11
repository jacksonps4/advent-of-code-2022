package main

type fileSystem struct {
	Root *treeNode
}

func NewFileSystem() *fileSystem {
	return &fileSystem{
		Root: &treeNode{
			Parent:   nil,
			Name:     "/",
			Files:    make(map[string]*file),
			Children: make(map[string]*treeNode),
		},
	}
}

func (fs *fileSystem) Size() int {
	return fs.Root.Size()
}

func (fs *fileSystem) FreeSpace() int {
	return 70000000 - fs.Size()
}

func (fs *fileSystem) GetPathNode(path []string) *treeNode {
	current := fs.Root
	for _, p := range path {
		child := current.Children[p]
		if child == nil {
			return nil
		}
		current = child
	}
	return current
}

func (fs *fileSystem) Walk(action func(node *treeNode)) {
	visit(fs.Root, action)
}

func visit(node *treeNode, action func(node *treeNode)) {
	if len(node.Children) > 0 {
		for _, child := range node.Children {
			visit(child, action)
		}
	}
	action(node)
}
