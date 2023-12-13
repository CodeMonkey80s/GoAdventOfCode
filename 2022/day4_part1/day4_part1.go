package day4_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type Range struct {
	Lo int
	Hi int
}

func pairsFullyContained(lines []string) int {
	sum := 0
	for _, line := range lines {
		ranges := extractRanges(line)
		if ranges[1].Lo >= ranges[0].Lo &&
			ranges[1].Hi <= ranges[0].Hi {
			sum++
			continue
		}
		if ranges[0].Lo >= ranges[1].Lo &&
			ranges[0].Hi <= ranges[1].Hi {
			sum++
			continue
		}
	}
	return sum
}

func extractRanges(line string) []Range {
	ranges := make([]Range, 0, 2)
	elves := strings.Split(line, ",")
	for _, elf := range elves {
		parts := strings.Split(elf, "-")
		r := Range{
			Lo: util.ConvertStringToInt(parts[0]),
			Hi: util.ConvertStringToInt(parts[1]),
		}
		ranges = append(ranges, r)
	}
	return ranges
}
