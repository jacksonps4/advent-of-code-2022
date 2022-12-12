package main

import (
	"errors"
)

type cpu struct {
	instructionList []*instruction
	programCounter  int
	lastInstruction *instruction
	pendingCycles   int
	xVal            int
	counter         int
}

func NewCPU() *cpu {
	return &cpu{
		xVal:    1,
		counter: 0,
	}
}

func (c *cpu) AddInstruction(instruction *instruction) {
	c.instructionList = append(c.instructionList, instruction)
}

func (c *cpu) Cycle() error {
	c.counter++
	if c.pendingCycles > 0 {
		c.pendingCycles--
		return nil
	}

	lastInstruction := c.lastInstruction
	if lastInstruction != nil {
		switch lastInstruction.opCode {
		case "addx":
			c.xVal += c.lastInstruction.arg
		}
		c.programCounter++
	}

	if c.programCounter >= len(c.instructionList) {
		return errors.New("no instructions left to execute")
	}

	nextInstruction := c.instructionList[c.programCounter]
	switch nextInstruction.opCode {
	case "addx":
		c.pendingCycles = 1
		break
	case "noop":
		c.pendingCycles = 0
		break
	}
	c.lastInstruction = nextInstruction

	return nil
}

func (c *cpu) SignalStrength() int {
	return c.xVal * c.counter
}

func (c *cpu) Count() int {
	return c.counter
}

func (c *cpu) ReadX() int {
	return c.xVal
}
