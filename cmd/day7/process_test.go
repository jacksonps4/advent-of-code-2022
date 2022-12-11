package main

import (
	"sort"
	"strings"
	"testing"
)

func TestProcess(t *testing.T) {
	fs := NewFileSystem()
	processShellData(strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`), fs)

	t.Run("dir e", func(t *testing.T) {
		fs.Walk(func(node *treeNode) {
			if node.Name == "e" {
				if node.Size() != 584 {
					t.Errorf("expected size %d but was %d", 584, node.Size())
				}
			}
		})
	})

	t.Run("dir a", func(t *testing.T) {
		fs.Walk(func(node *treeNode) {
			if node.Name == "a" {
				if node.Size() != 94853 {
					t.Errorf("expected size %d but was %d", 94853, node.Size())
				}
			}
		})
	})

	t.Run("dir d", func(t *testing.T) {
		fs.Walk(func(node *treeNode) {
			if node.Name == "d" {
				if node.Size() != 24933642 {
					t.Errorf("expected size 24933642 but was %d", node.Size())
				}
			}
		})
	})

	t.Run("dir /", func(t *testing.T) {
		fs.Walk(func(node *treeNode) {
			if node.Name == "/" {
				if node.Size() != 48381165 {
					t.Errorf("expected size 48381165 but was %d", node.Size())
				}
			}
		})
	})

	t.Run("total size", func(t *testing.T) {
		sum := 0
		fs.Walk(func(node *treeNode) {
			if node.Size() <= 100000 {
				sum += node.Size()
			}
		})
		if sum != 95437 {
			t.Errorf("expected 95437 but was %d", sum)
		}
	})

	t.Run("Free space", func(t *testing.T) {
		freeSpace := fs.FreeSpace()
		spaceRequired := 30000000 - freeSpace
		var bigEnough []*treeNode
		fs.Walk(func(node *treeNode) {
			sz := node.Size()
			if sz > spaceRequired {
				bigEnough = append(bigEnough, node)
			}
		})
		sort.SliceStable(bigEnough, func(i, j int) bool {
			return bigEnough[i].Size() < bigEnough[j].Size()
		})
		maxSz := bigEnough[0].Size()
		if maxSz != 24933642 {
			t.Errorf("expected 24933642 but was %d", maxSz)
		}
	})
}

func TestProcessDots(t *testing.T) {
	result := processDots("/foo/bar/test/..")
	if result != "/foo/bar" {
		t.Errorf("expected %s but was %s", "/foo/bar/test", result)
	}
}
