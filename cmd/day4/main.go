package main

import (
	"bufio"
	"fmt"
	"io"
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

	ranges := parseSectionRanges(file)
	containedCount, overlapCount := checkPairs(ranges)

	fmt.Printf("Part 1: found %d fully contained regions\n", containedCount)
	fmt.Printf("Part 2: found %d overlapping regions\n", overlapCount)
}

func parseSectionRanges(file io.Reader) []*sectionRangePair {
	scanner := bufio.NewScanner(file)
	ranges := make([]*sectionRangePair, 0)
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
		if len(first) != 2 {
			panic(fmt.Sprintf("invalid data in input file: found multiple pairs"))
		}

		second := strings.Split(values[1], "-")
		if len(second) != 2 {
			panic(fmt.Sprintf("invalid data in input file: found multiple pairs"))
		}

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
	return ranges
}

func checkPairs(ranges []*sectionRangePair) (int, int) {
	containedCount := 0
	overlapCount := 0

	for _, r := range ranges {
		if r.isContained() {
			containedCount++
		}
		if r.overlaps() {
			overlapCount++
		}
	}

	return containedCount, overlapCount
}
