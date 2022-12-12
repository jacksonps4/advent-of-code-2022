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
	defer file.Close()

	cpu := NewCPU()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := ParseInstruction(scanner.Text())
		cpu.AddInstruction(instruction)
	}

	sum := 0
	crt := NewCRT()
	for {
		err := cpu.Cycle()
		if err != nil {
			break
		}

		count := cpu.Count()
		if (count-20)%40 == 0 {
			sum += cpu.SignalStrength()
		}

		spriteHPos := cpu.xVal
		crtPosition := crt.Position() % 40
		if crtPosition == spriteHPos || (spriteHPos+1) == crtPosition || (spriteHPos-1) == crtPosition {
			crt.Set()
		} else {
			crt.Unset()
		}
		crt.Cycle()
	}

	fmt.Printf("Part 1: sum = %d\n\n", sum)
	fmt.Printf("%s", crt)
}
