package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}

	stream, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("failed to read buffer: %s", err))
	}

	v := processStream(stream, 4)
	if v > 0 {
		fmt.Printf("Part 1: %v\n", v)
	} else {
		panic(fmt.Sprintf("Invalid result: %d", v))
	}

	v = processStream(stream, 14)
	if v > 0 {
		fmt.Printf("Part 2: %v\n", v)
	} else {
		panic(fmt.Sprintf("Invalid result: %d", v))
	}

}

func processStream(stream []byte, l int) int {
	buf := make([]byte, l)
	for i := 0; i < len(stream); i++ {
		index := i % l
		buf[index] = stream[i]
		if i < l {
			continue
		}
		if checkForUniqueness(buf) {
			return i + 1
		}
	}

	return -1
}

func checkForUniqueness(buf []byte) bool {
	counts := make(map[byte]int)
	for _, b := range buf {
		counts[b]++
	}
	return len(counts) == len(buf)
}
