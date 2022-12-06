package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	elfCalories := make(map[int]int)
	currentCalories := 0
	for i := 0; scanner.Scan(); {
		line := scanner.Text()
		if line == "" {
			elfCalories[i] = currentCalories
			currentCalories = 0
			i++
			continue
		}
		value, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("failed to parse line: %s; error = %s", line, err))
		}
		currentCalories += value
	}

	keys := make([]int, len(elfCalories))
	for key := range elfCalories {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return elfCalories[keys[i]] >= elfCalories[keys[j]]
	})

	fmt.Printf("Part 1: the elf in position %d has %d calories\n", keys[0], elfCalories[keys[0]])

	var calories int
	for i := 0; i < 3; i++ {
		calories += elfCalories[keys[i]]
	}
	fmt.Printf("Part 2: the top three elves have %d calories\n", calories)
}
