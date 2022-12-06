package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	part1Points := 0
	part2Points := 0
	for scanner.Scan() {
		line := scanner.Text()

		columns := strings.Split(line, " ")

		opponentChoses := columns[0]
		suggestedPlay := columns[1]

		yourScorePart1, _ := calculateScorePart1(suggestedPlay, opponentChoses)
		_, yourScorePart2 := calculateScorePart2(opponentChoses, suggestedPlay)

		part1Points += yourScorePart1
		part2Points += yourScorePart2
	}

	fmt.Printf("PART 1: score = %d\n", part1Points)
	fmt.Printf("PART 2: score = %d\n", part2Points)
}

func calculateScorePart1(left string, right string) (int, int) {
	l := ParseRockPaperScissorsString(left)
	r := ParseRockPaperScissorsString(right)
	result := l.WinsAgainst(r)
	if result == nil {
		// draw
		return l.Points() + 3, r.Points() + 3
	} else {
		switch *result {
		case true:
			return l.Points() + 6, r.Points()
		case false:
			return l.Points(), r.Points() + 6
		}
	}

	panic("no result")
}

func calculateScorePart2(left string, strategy string) (int, int) {
	l := ParseRockPaperScissorsString(left)
	var r rockPaperScissors
	switch strategy {
	case "X":
		// need to lose
		r = l.Beats()
	case "Y":
		// need to draw
		r = l
	case "Z":
		// need to win
		r = l.IsBeatenBy()
	}

	result := l.WinsAgainst(r)
	if result == nil {
		// draw
		return l.Points() + 3, r.Points() + 3
	} else {
		switch *result {
		case true:
			return l.Points() + 6, r.Points()
		case false:
			return l.Points(), r.Points() + 6
		}
	}

	panic("no result")
}
