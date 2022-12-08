package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	b := bytes.Buffer{}
	_, err = io.Copy(&b, file)
	if err != nil {
		panic(fmt.Sprintf("can't copy file data: %s", err))
	}

	scanner := bufio.NewScanner(strings.NewReader(b.String()))
	maxX := 0
	maxY := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		line := scanner.Text()
		lineLength := len(line)
		if lineLength > maxX {
			maxX = lineLength
		}
		maxY++
	}

	forest := NewForest(maxX, maxY)
	scanner = bufio.NewScanner(strings.NewReader(b.String()))
	y := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		line := scanner.Text()
		for x, c := range line {
			forest.SetHeight(x, y, int(c)-48)
		}
		y++
	}

	visible := 0
	maxScenicScore := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if forest.IsVisible(x, y) {
				visible++
			}

			scenicScore := forest.ScenicScore(x, y)
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	fmt.Printf("Part 1: %d trees are visible\n", visible)
	fmt.Printf("Part 2: Max scenic score is %d\n", maxScenicScore)
}
