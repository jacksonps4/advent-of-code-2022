package main

import "fmt"

type rockPaperScissors rune

const (
	ROCK     rockPaperScissors = 'A'
	PAPER    rockPaperScissors = 'B'
	SCISSORS rockPaperScissors = 'C'
)

func ParseRockPaperScissorsString(v string) rockPaperScissors {
	if len(v) > 1 || len(v) == 0 {
		panic(fmt.Sprintf("invalid RPS value: %s", v))
	}
	switch v {
	case "A":
		fallthrough
	case "X":
		return ROCK
	case "B":
		fallthrough
	case "Y":
		return PAPER
	case "C":
		fallthrough
	case "Z":
		return SCISSORS
	default:
		panic(fmt.Sprintf("invalid RPS value: %s", v))
	}
}

func (left rockPaperScissors) Points() int {
	switch left {
	case ROCK:
		return 1
	case PAPER:
		return 2
	case SCISSORS:
		return 3
	}
	return -1
}

func (left rockPaperScissors) Beats() rockPaperScissors {
	switch left {
	case ROCK:
		return SCISSORS
	case PAPER:
		return ROCK
	case SCISSORS:
		return PAPER
	}

	panic("invalid input")
}

func (left rockPaperScissors) IsBeatenBy() rockPaperScissors {
	switch left {
	case ROCK:
		return PAPER
	case PAPER:
		return SCISSORS
	case SCISSORS:
		return ROCK
	}

	panic("invalid input")
}

func (left rockPaperScissors) WinsAgainst(right rockPaperScissors) *bool {
	win := true
	loss := false
	var draw *bool

	switch left {
	case ROCK:
		switch right {
		case PAPER:
			return &loss
		case SCISSORS:
			return &win
		}
	case PAPER:
		switch right {
		case ROCK:
			return &win
		case SCISSORS:
			return &loss
		}
	case SCISSORS:
		switch right {
		case ROCK:
			return &loss
		case PAPER:
			return &win
		}
	}

	return draw
}
