package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read stack structure
	inputs := NewStack()
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		data := scanner.Text()
		if data == "" {
			break
		}
		inputs.Push(data)
	}

	var stacks []*stack
	for i := 0; i < 9; i++ {
		stacks = append(stacks, NewStack())
	}

	for {
		data := inputs.Pop()
		if data == "" {
			break
		}

		// 1, 5, 9, 13, 17, 21, 25, 29, 33
		index := 0
		for i := 1; i < len(data); i += 4 {
			item := string(data[i])
			if item != " " {
				stacks[index].Push(item)
			}
			index++
		}
	}

	// read movements
	movementRe := regexp.MustCompile("move ([0-9][0-9]*) from ([0-9]) to ([0-9])")
	movements := make([]*movement, 0)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		data := scanner.Text()

		movementData := movementRe.ReplaceAllString(data, "$1 $2 $3")
		elements := strings.Split(movementData, " ")
		movements = append(movements, &movement{
			Count: mustParseInt(elements[0]),
			From:  mustParseInt(elements[1]),
			To:    mustParseInt(elements[2]),
		})
	}

	// execute the movements
	//executeMovementsPart1(movements, stacks)
	executeMovementsPart2(movements, stacks)

	for _, stack := range stacks {
		fmt.Printf("%s", stack.Pop())
	}
	fmt.Printf("\n")
}

func executeMovementsPart1(movements []*movement, stacks []*stack) {
	for _, movement := range movements {
		for i := 0; i < movement.Count; i++ {
			retrieved := stacks[movement.From-1].Pop()
			stacks[movement.To-1].Push(retrieved)
		}
	}
}

func executeMovementsPart2(movements []*movement, stacks []*stack) {
	for _, movement := range movements {
		toBeMoved := make([]string, 0)
		for i := 0; i < movement.Count; i++ {
			retrieved := stacks[movement.From-1].Pop()
			toBeMoved = append(toBeMoved, retrieved)
		}
		for i := movement.Count - 1; i >= 0; i-- {
			stacks[movement.To-1].Push(toBeMoved[i])
		}
	}
}

func mustParseInt(data string) int {
	i, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}
	return i
}
