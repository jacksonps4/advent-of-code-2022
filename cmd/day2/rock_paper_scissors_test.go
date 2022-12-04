package main

import "testing"

func TestRps_WinsAgainst(t *testing.T) {
	win := true
	loss := false

	scenarios := []struct {
		Name           string
		Left           rockPaperScissors
		Right          rockPaperScissors
		ExpectedResult *bool
	}{
		{
			Name:           "Rock blunts scissors",
			Left:           ROCK,
			Right:          SCISSORS,
			ExpectedResult: &win,
		},
		{
			Name:           "Paper wraps rock",
			Left:           PAPER,
			Right:          ROCK,
			ExpectedResult: &win,
		},
		{
			Name:           "Scissors cut paper",
			Left:           SCISSORS,
			Right:          PAPER,
			ExpectedResult: &win,
		},
		{
			Name:           "Scissors blunted by rock",
			Left:           SCISSORS,
			Right:          ROCK,
			ExpectedResult: &loss,
		},
		{
			Name:           "Rock wrapped by paper",
			Left:           ROCK,
			Right:          PAPER,
			ExpectedResult: &loss,
		},
		{
			Name:           "Paper cut by scissors",
			Left:           PAPER,
			Right:          SCISSORS,
			ExpectedResult: &loss,
		},
		{
			Name:           "Rock draws with rock",
			Left:           ROCK,
			Right:          ROCK,
			ExpectedResult: nil,
		},
		{
			Name:           "Scissors draws with scissors",
			Left:           SCISSORS,
			Right:          SCISSORS,
			ExpectedResult: nil,
		},
		{
			Name:           "Paper draws with paper",
			Left:           PAPER,
			Right:          PAPER,
			ExpectedResult: nil,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			result := scenario.Left.WinsAgainst(scenario.Right)
			if (result == nil && scenario.ExpectedResult != nil) || (result != nil && scenario.ExpectedResult == nil) {
				t.Errorf("Scenario %s: expected %v but was %v", scenario.Name, *scenario.ExpectedResult, *result)
			}
			if result != nil && scenario.ExpectedResult != nil && *result != *scenario.ExpectedResult {
				t.Errorf("Scenario %s: expected %v but was %v", scenario.Name, *scenario.ExpectedResult, *result)
			}
		})
	}
}
