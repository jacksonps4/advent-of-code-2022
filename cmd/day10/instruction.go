package main

import (
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	opCode string
	arg    int
}

func ParseInstruction(i string) *instruction {
	instr := strings.Split(i, " ")
	opCode := instr[0]
	argument := 0
	if len(instr) > 1 {
		arg, err := strconv.Atoi(instr[1])
		if err != nil {
			panic(fmt.Sprintf("failed to read argument: %s", err))
		}
		argument = arg
	}

	return &instruction{
		opCode: opCode,
		arg:    argument,
	}
}

func (i *instruction) String() string {
	return fmt.Sprintf("%s %d", i.opCode, i.arg)
}
