package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}

	scanner := bufio.NewScanner(file)
	priorities := 0
	count := 0
	group := make([]string, 0)
	groups := make([][]string, 0)
	for scanner.Scan() {
		data := scanner.Text()
		priorities += getPrioritySum(data)
		if count > 0 && count%3 == 0 {
			groups = append(groups, group)
			group = make([]string, 0)
			group = append(group, data)
		} else {
			group = append(group, data)
		}
		count++
	}
	if len(group) > 0 {
		groups = append(groups, group)
	}

	fmt.Printf("Part 1: priorities = %d\n", priorities)

	groupPriority := 0
	for _, group := range groups {
		intersections := getIntersectionStrings(group)
		if len(intersections) == 1 {
			groupPriority += getPriority(intersections[0])
		} else {
			panic(fmt.Sprintf("more than one intersection"))
		}
	}

	fmt.Printf("Part 2: priorities = %d\n", groupPriority)
}

func getPrioritySum(itemList string) int {
	priorities := 0

	intersections := getIntersections(itemList)

	for _, v := range intersections {
		priorities += getPriority(v)
	}

	return priorities
}

func getIntersections(itemList string) []rune {
	middle := (len(itemList) / 2)

	firstHalf := itemList[0:middle]
	secondHalf := itemList[middle:]

	return getIntersectionStrings([]string{firstHalf, secondHalf})
}

func getIntersectionStrings(values []string) []rune {
	indexes := make(map[rune]int)

	for _, v := range values {
		currentSet := make(map[rune]int)
		for _, r := range v {
			currentSet[r]++
		}
		for r, _ := range currentSet {
			indexes[r]++
		}
	}

	size := len(values)
	result := make([]rune, 0)
	for k, v := range indexes {
		if v == size {
			result = append(result, k)
		}
	}

	return result
}

func getPriority(v rune) int {
	if v >= 65 && v <= 90 {
		// uppercase: 27-52
		return int(v - 38)
	} else if v >= 97 && v <= 122 {
		// lowercase: 1-26
		return int(v - 96)
	} else {
		panic(fmt.Sprintf("invalid value: %v", v))
	}
}
