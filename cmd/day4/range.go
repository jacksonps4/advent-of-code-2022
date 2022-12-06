package main

import "fmt"

type sectionRangePair struct {
	Start1 int
	End1   int
	Start2 int
	End2   int
}

func (r *sectionRangePair) String() string {
	return fmt.Sprintf("%d-%d,%d-%d", r.Start1, r.End1, r.Start2, r.End2)
}
