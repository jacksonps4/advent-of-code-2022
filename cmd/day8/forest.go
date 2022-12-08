package main

type forest struct {
	Heights [][]int
}

func NewForest(xmax int, ymax int) *forest {
	f := &forest{
		Heights: make([][]int, ymax),
	}
	for i, _ := range f.Heights {
		f.Heights[i] = make([]int, xmax)
	}
	return f
}

func (f *forest) GetHeight(x int, y int) int {
	return f.Heights[y][x]
}

func (f *forest) SetHeight(x int, y int, height int) {
	f.Heights[y][x] = height
}

func (f *forest) IsVisible(x int, y int) bool {
	return f.isVisibleRight(x, y) || f.isVisibleTop(x, y) ||
		f.isVisibleLeft(x, y) || f.isVisibleBottom(x, y)
}

// 49142
// 12345
// 12345
// 53262
//
// x=3, y=1 (3,1)

func (f *forest) ScenicScore(x int, y int) int {
	treeHeight := f.GetHeight(x, y)

	// look to the top
	topScore := 0
	for v := y - 1; v >= 0; v-- {
		topScore++
		h := f.GetHeight(x, v)
		if h >= treeHeight {
			break
		}
	}

	// look to the bottom
	bottomScore := 0
	for v := y + 1; v < len(f.Heights); v++ {
		bottomScore++
		h := f.GetHeight(x, v)
		if h >= treeHeight {
			break
		}
	}

	// look to the right
	rightScore := 0
	for v := x + 1; v < len(f.Heights); v++ {
		rightScore++
		h := f.GetHeight(v, y)
		if h >= treeHeight {
			break
		}
	}

	// look to the left
	leftScore := 0
	for v := x - 1; v >= 0; v-- {
		leftScore++
		h := f.GetHeight(v, y)
		if h >= treeHeight {
			break
		}
	}

	return leftScore * rightScore * bottomScore * topScore
}

func (f *forest) isVisibleTop(x int, y int) bool {
	treeHeight := f.GetHeight(x, y)
	for v := y - 1; v >= 0; v-- {
		h := f.GetHeight(x, v)
		if h >= treeHeight {
			return false
		}
	}
	return true
}

func (f *forest) isVisibleBottom(x int, y int) bool {
	treeHeight := f.GetHeight(x, y)
	for v := y + 1; v < len(f.Heights); v++ {
		h := f.GetHeight(x, v)
		if h >= treeHeight {
			return false
		}
	}
	return true

}

func (f *forest) isVisibleRight(x int, y int) bool {
	treeHeight := f.GetHeight(x, y)
	for v := x + 1; v < len(f.Heights); v++ {
		h := f.GetHeight(v, y)
		if h >= treeHeight {
			return false
		}
	}
	return true
}

func (f *forest) isVisibleLeft(x int, y int) bool {
	treeHeight := f.GetHeight(x, y)
	for v := x - 1; v >= 0; v-- {
		h := f.GetHeight(v, y)
		if h >= treeHeight {
			return false
		}
	}
	return true
}
