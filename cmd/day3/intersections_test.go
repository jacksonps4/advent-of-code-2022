package main

import "testing"

func TestPrioritySums(t *testing.T) {
	scenarios := []struct {
		Name         string
		ItemList     string
		Intersection rune
		ExpectedSum  int
	}{
		{
			Name:         "first",
			ItemList:     "vJrwpWtwJgWrhcsFMMfFFhFp",
			Intersection: 'p',
			ExpectedSum:  16,
		},
		{
			Name:         "second",
			ItemList:     "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			Intersection: 'L',
			ExpectedSum:  38,
		},
		{
			Name:         "third",
			ItemList:     "PmmdzqPrVvPwwTWBwg",
			Intersection: 'P',
			ExpectedSum:  42,
		},
		{
			Name:         "fourth",
			ItemList:     "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			Intersection: 'v',
			ExpectedSum:  22,
		},
		{
			Name:         "fifth",
			ItemList:     "ttgJtRGJQctTZtZT",
			Intersection: 't',
			ExpectedSum:  20,
		},
		{
			Name:         "sixth",
			ItemList:     "CrZsJsPPZsGzwwsLwLmpwMDw",
			Intersection: 's',
			ExpectedSum:  19,
		},
	}

	for _, scenario := range scenarios {
		t.Run("intersection "+scenario.Name, func(t *testing.T) {
			v := getIntersections(scenario.ItemList)
			expected := scenario.Intersection
			if v[0] != expected {
				t.Errorf("expected %s but was %s", string(expected), string(v))
			}
		})
		t.Run("sum "+scenario.Name, func(t *testing.T) {
			v := getPrioritySum(scenario.ItemList)
			expected := scenario.ExpectedSum
			if v != expected {
				t.Errorf("expected %d but was %d", expected, v)
			}
		})
	}
}

func TestBadges(t *testing.T) {
	scenarios := []struct {
		Name     string
		Input    []string
		Expected []rune
		Priority int
	}{
		{
			Name: "First group",
			Input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
			},
			Expected: []rune{'r'},
			Priority: 18,
		},
		{
			Name: "Second group",
			Input: []string{
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			Expected: []rune{'Z'},
			Priority: 52,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.Name, func(t *testing.T) {
			intersections := getIntersectionStrings(scenario.Input)
			for i, v := range intersections {
				priority := getPriority(v)
				if scenario.Expected[i] != v {
					t.Errorf("Expected %v but was %v", scenario.Expected[i], v)
				}
				if scenario.Priority != priority {
					t.Errorf("Expected priority %d but was %d", scenario.Priority, priority)
				}
			}
		})
	}
}
