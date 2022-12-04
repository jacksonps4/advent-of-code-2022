package main

import "testing"

func TestStrategy1(t *testing.T) {
	t.Run("Round 1", func(t *testing.T) {
		_, yourScore := calculateScorePart1("A", "Y")
		if yourScore != 8 {
			t.Errorf("Expected 8 but was %d", yourScore)
		}
	})
	t.Run("Round 2", func(t *testing.T) {
		_, yourScore := calculateScorePart1("B", "X")
		if yourScore != 1 {
			t.Errorf("Expected your score = 1 but was %d", yourScore)
		}
	})
	t.Run("Round 3", func(t *testing.T) {
		yourScore, opponentsScore := calculateScorePart1("C", "Z")
		if yourScore != 6 {
			t.Errorf("Expected 6 but was %d", yourScore)
		}
		if opponentsScore != 6 {
			t.Errorf("Expected 6 but was %d", opponentsScore)
		}
	})
}

func TestStrategy2(t *testing.T) {
	t.Run("Round 1", func(t *testing.T) {
		opponentsScore, yourScore := calculateScorePart2("A", "Y")
		if yourScore != 4 {
			t.Errorf("Expected 4 but was %d", yourScore)
		}
		if opponentsScore != 4 {
			t.Errorf("Expected 4 but was %d", yourScore)
		}
	})
	t.Run("Round 2", func(t *testing.T) {
		_, yourScore := calculateScorePart2("B", "X")
		if yourScore != 1 {
			t.Errorf("Expected your score = 1 but was %d", yourScore)
		}
	})
	t.Run("Round 3", func(t *testing.T) {
		_, yourScore := calculateScorePart2("C", "Z")
		if yourScore != 7 {
			t.Errorf("Expected 7 but was %d", yourScore)
		}
	})
}
