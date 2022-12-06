package main

import (
	"fmt"
	"testing"
)

func TestStreamProcess(t *testing.T) {
	scenarios := []struct {
		Input    []byte
		Position int
	}{
		{
			Input:    []byte(`bvwbjplbgvbhsrlpgdmjqwftvncz`),
			Position: 5,
		},
		{
			Input:    []byte(`nppdvjthqldpwncqszvftbrmjlhg`),
			Position: 6,
		},
		{
			Input:    []byte(`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`),
			Position: 10,
		},
		{
			Input:    []byte(`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`),
			Position: 11,
		},
	}

	for i, scenario := range scenarios {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			position := processStream(scenario.Input)
			if position != scenario.Position {
				t.Errorf("expected position %d but was %d", scenario.Position, position)
			}
		})
	}
}
