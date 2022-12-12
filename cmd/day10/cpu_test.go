package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestCPU1(t *testing.T) {
	instructionSet := `noop
addx 3
addx -5`

	cpu := NewCPU()
	scanner := bufio.NewScanner(strings.NewReader(instructionSet))

	for scanner.Scan() {
		instruction := ParseInstruction(scanner.Text())
		cpu.AddInstruction(instruction)
	}

	// initial state
	x := cpu.ReadX()
	if x != 1 {
		fmt.Errorf("expected X = 1 but was %d", x)
	}

	// 1st cycle
	cpu.Cycle()
	x = cpu.ReadX()
	if x != 1 {
		fmt.Errorf("expected X = 1 but was %d", x)
	}

	// 2nd cycle
	cpu.Cycle()
	x = cpu.ReadX()
	if x != 1 {
		fmt.Errorf("expected X = 1 but was %d", x)
	}

	// 3rd cycle
	cpu.Cycle()
	x = cpu.ReadX()
	if x != 4 {
		fmt.Errorf("expected X = 4 but was %d", x)
	}

	// 4th cycle
	cpu.Cycle()
	x = cpu.ReadX()
	if x != 4 {
		fmt.Errorf("expected X = 4 but was %d", x)
	}

	// 5th cycle
	cpu.Cycle()
	x = cpu.ReadX()
	if x != -1 {
		fmt.Errorf("expected X = -1 but was %d", x)
	}

	// 6th cycle
	err := cpu.Cycle()
	if err == nil {
		fmt.Errorf("expected end of instruction set")
	}
	if x != -1 {
		fmt.Errorf("expected X = -1 but was %d", x)
	}
}

func TestCPU2(t *testing.T) {
	instructionSet := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
`
	cpu := NewCPU()
	scanner := bufio.NewScanner(strings.NewReader(instructionSet))

	for scanner.Scan() {
		instruction := ParseInstruction(scanner.Text())
		cpu.AddInstruction(instruction)
	}

	sum := 0
	for {
		err := cpu.Cycle()
		if err != nil {
			break
		}
		count := cpu.Count()
		if (count-20)%40 == 0 {
			strength := cpu.SignalStrength()
			sum += strength
		}
	}

	if sum != 13140 {
		t.Errorf("expected %d but was %d", 13140, sum)
	}
}
