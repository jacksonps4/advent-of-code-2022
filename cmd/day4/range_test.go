package main

import (
	"fmt"
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
			contained := isContained(s.Pair)
			if s.Contained != contained {
				t.Errorf("Expected %v but was %v", s.Contained, contained)
			}
		})
		t.Run(fmt.Sprintf("overlaps pair %d", i+1), func(t *testing.T) {
			overlaps := overlaps(s.Pair)
			if s.Overlaps != overlaps {
				t.Errorf("Expected %v but was %v", s.Overlaps, overlaps)
			}
		})
	}
}
