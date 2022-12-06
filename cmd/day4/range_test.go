package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRangeExample(t *testing.T) {
	scenarios := []struct {
		Pair      *sectionRangePair
		Contained bool
		Overlaps  bool
	}{
		{
			Pair: &sectionRangePair{
				Start1: 2,
				End1:   4,
				Start2: 6,
				End2:   8,
			},
			Contained: false,
			Overlaps:  false,
		},
		{
			Pair: &sectionRangePair{
				Start1: 2,
				End1:   3,
				Start2: 4,
				End2:   5,
			},
			Contained: false,
			Overlaps:  false,
		},
		{
			Pair: &sectionRangePair{
				Start1: 5,
				End1:   7,
				Start2: 7,
				End2:   9,
			},
			Contained: false,
			Overlaps:  true,
		},
		{
			Pair: &sectionRangePair{
				Start1: 2,
				End1:   8,
				Start2: 3,
				End2:   7,
			},
			Contained: true,
			Overlaps:  true,
		},
		{
			Pair: &sectionRangePair{
				Start1: 6,
				End1:   6,
				Start2: 4,
				End2:   6,
			},
			Contained: true,
			Overlaps:  true,
		},
		{
			Pair: &sectionRangePair{
				Start1: 2,
				End1:   6,
				Start2: 4,
				End2:   8,
			},
			Contained: false,
			Overlaps:  true,
		},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("contained pair %d", i+1), func(t *testing.T) {
			contained := s.Pair.isContained()
			if s.Contained != contained {
				t.Errorf("Expected %v but was %v", s.Contained, contained)
			}
		})
		t.Run(fmt.Sprintf("overlaps pair %d", i+1), func(t *testing.T) {
			overlaps := s.Pair.overlaps()
			if s.Overlaps != overlaps {
				t.Errorf("Expected %v but was %v", s.Overlaps, overlaps)
			}
		})
	}
}

func TestInput(t *testing.T) {
	rawData := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	data := strings.NewReader(rawData)
	ranges := parseSectionRanges(data)

	containedCount, overlapCount := checkPairs(ranges)
	if containedCount != 2 {
		t.Errorf("expected 2 but was %d", containedCount)
	}

	if overlapCount != 4 {
		t.Errorf("expected 4 but was %d", overlapCount)
	}
}
