package day5_part1

import (
	"GoAdventOfCode/util"
	"strings"
)

func getAnswerForPart1(lines []string) int {

	var fresh int

	ranges, idx := loadRanges(lines)
	values := loadValues(lines[idx+1:])

loop:
	for _, v := range values {
		for _, r := range ranges {
			if v >= r.MinVal && v <= r.MaxVal {
				fresh++
				continue loop
			}
		}
	}

	return fresh
}

type valueRange struct {
	MinVal int
	MaxVal int
}

func loadRanges(lines []string) ([]valueRange, int) {

	var ranges []valueRange
	for i, line := range lines {
		if len(line) == 0 {
			return ranges, i
		}
		if strings.IndexByte(line, '-') != -1 {
			parts := strings.Split(line, "-")

			r := valueRange{
				MinVal: util.ConvertStringToInt(parts[0]),
				MaxVal: util.ConvertStringToInt(parts[1]),
			}
			ranges = append(ranges, r)
		}
	}
	return ranges, 0
}

func loadValues(lines []string) []int {
	var values []int
	for _, num := range lines {
		values = append(values, util.ConvertStringToInt(num))
	}

	return values
}
