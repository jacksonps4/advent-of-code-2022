package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestForest(t *testing.T) {
	data := `30373
25512
65332
33549
35390`
	forest := NewForest(5, 5)

	scanner := bufio.NewScanner(strings.NewReader(data))
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			forest.SetHeight(x, y, int(c)-48)
		}
		y++
	}

	t.Run("Scenario 1", func(t *testing.T) {
		if !forest.isVisibleTop(1, 1) {
			t.Errorf("Needs to be visible from top")
		}
		if !forest.isVisibleLeft(1, 1) {
			t.Errorf("Needs to be visible from left")
		}
		if forest.isVisibleBottom(1, 1) {
			t.Errorf("Needs to not be visible from bottom")
		}
		if forest.isVisibleRight(1, 1) {
			t.Errorf("Needs to not be visible from right")
		}
	})
	t.Run("Scenario 2", func(t *testing.T) {
		if !forest.isVisibleTop(2, 1) {
			t.Errorf("Needs to be visible from top")
		}
		if forest.isVisibleLeft(2, 1) {
			t.Errorf("Needs to not be visible from left")
		}
		if forest.isVisibleBottom(2, 1) {
			t.Errorf("Needs to not be visible from bottom")
		}
		if !forest.isVisibleRight(2, 1) {
			t.Errorf("Needs to be visible from right")
		}
	})
	t.Run("Scenario 3", func(t *testing.T) {
		if forest.IsVisible(3, 1) {
			t.Errorf("Should not be visible")
		}
	})
	t.Run("Scenario 4", func(t *testing.T) {
		if forest.isVisibleTop(1, 2) {
			t.Errorf("Needs to not be visible from top")
		}
		if forest.isVisibleLeft(1, 2) {
			t.Errorf("Needs to not be visible from left")
		}
		if forest.isVisibleBottom(1, 2) {
			t.Errorf("Needs to not be visible from bottom")
		}
		if !forest.isVisibleRight(1, 2) {
			t.Errorf("Needs to be visible from right")
		}
	})
	t.Run("Scenario 5", func(t *testing.T) {
		if forest.IsVisible(2, 2) {
			t.Errorf("Should not be visible")
		}
	})
	t.Run("Scenario 6", func(t *testing.T) {
		if forest.isVisibleTop(3, 2) {
			t.Errorf("Needs to not be visible from top")
		}
		if forest.isVisibleLeft(3, 2) {
			t.Errorf("Needs to not be visible from left")
		}
		if forest.isVisibleBottom(3, 2) {
			t.Errorf("Needs to not be visible from bottom")
		}
		if !forest.isVisibleRight(3, 2) {
			t.Errorf("Needs to be visible from right")
		}
	})
	t.Run("Scenario 7", func(t *testing.T) {
		if !forest.IsVisible(2, 3) {
			t.Errorf("Needs to be visible")
		}
		if forest.IsVisible(1, 3) {
			t.Errorf("Needs to not be visible")
		}
		if forest.IsVisible(3, 3) {
			t.Errorf("Needs to not be visible")
		}
	})

	t.Run("Edges are visible", func(t *testing.T) {
		for y := 0; y < 5; y++ {
			if !forest.IsVisible(0, y) {
				t.Errorf("Edges need to be visible")
			}
		}
		for y := 0; y < 5; y++ {
			if !forest.IsVisible(4, y) {
				t.Errorf("Edges need to be visible")
			}
		}
		for x := 0; y < 5; x++ {
			if !forest.IsVisible(x, 0) {
				t.Errorf("Edges need to be visible")
			}
		}
		for x := 0; y < 5; x++ {
			if !forest.IsVisible(x, 4) {
				t.Errorf("Edges need to be visible")
			}
		}
	})

	t.Run("Scenic score 1", func(t *testing.T) {
		scenicScore := forest.ScenicScore(2, 1)
		if scenicScore != 4 {
			t.Errorf("Expected 4 but was %d", scenicScore)
		}
	})
	t.Run("Scenic score 2", func(t *testing.T) {
		scenicScore := forest.ScenicScore(2, 3)
		if scenicScore != 8 {
			t.Errorf("Expected 8 but was %d", scenicScore)
		}
	})
}
