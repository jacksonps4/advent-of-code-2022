package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Getenv("INPUT_FILE"))
	if err != nil {
		panic(fmt.Sprintf("can't read file: %s", err))
	}
	defer file.Close()

	ranges := make([]*sectionRangePair, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			panic(fmt.Sprintf("error reading: %s", scanner.Err()))
		}
		data := scanner.Text()
		values := strings.Split(data, ",")
		if len(values) != 2 {
			panic(fmt.Sprintf("invalid data in input file: found multiple pairs"))
		}

		first := strings.Split(values[0], "-")
		second := strings.Split(values[1], "-")

		start1, err := strconv.Atoi(first[0])
		if err != nil {
			panic(fmt.Sprintf("invalid data in input file: %s", err))
		}
		end1, err := strconv.Atoi(first[1])
		if err != nil {
			panic(fmt.Sprintf("invalid data in input file: %s", err))
		}
		start2, err := strconv.Atoi(second[0])
		if err != nil {
			panic(fmt.Sprintf("invalid data in input file: %s", err))
		}
		end2, err := strconv.Atoi(second[1])
		if err != nil {
			panic(fmt.Sprintf("invalid data in input file: %s", err))
		}

		ranges = append(ranges, &sectionRangePair{
			Start1: start1,
			End1:   end1,
			Start2: start2,
			End2:   end2,
		})
	}

	containedCount := findFullyContainedRegions(ranges)
	fmt.Printf("Part 1: found %d fully contained regions\n", containedCount)

	overlappingCount := findOverlappingRegions(ranges)
	fmt.Printf("Part 2: found %d overlapping regions\n", overlappingCount)
}

func findFullyContainedRegions(ranges []*sectionRangePair) int {
	count := 0
	for _, r := range ranges {
		if isContained(r) {
			count++
		}
	}
	return count
}

func findOverlappingRegions(ranges []*sectionRangePair) int {
	count := 0
	for _, r := range ranges {
		if overlaps(r) {
			count++
			//fmt.Printf("%s\n", r)
		}
	}
	return count
}

func isContained(r *sectionRangePair) bool {
	return (r.Start1 <= r.Start2 && r.End1 >= r.End2) || (r.Start2 <= r.Start1 && r.End2 >= r.End1)
}

func overlaps(r *sectionRangePair) bool {
	return r.Start2 <= r.End1 && r.End1 >= r.Start2
}
