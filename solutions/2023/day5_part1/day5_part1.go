package day5_part1

import (
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

type Range struct {
	Dst    int
	Src    int
	Length int
}

func lowestLocationNumber(lines []string) int {
	seeds := extractSeeds(lines[0])
	almanac := extractAlmanac(lines[1:])
	return processSeedsWithAlmanac(seeds, almanac)
}

func processSeedsWithAlmanac(seeds []int, groups [][]Range) int {
	locations := make([]int, 0)
	for _, seed := range seeds {
		value := seed
		for _, ranges := range groups {
			for _, r := range ranges {
				srcStart := r.Src
				srcEnd := r.Src + r.Length
				dstStart := r.Dst
				if value >= srcStart && value <= srcEnd {
					value = dstStart + value - srcStart
					break
				}
			}
		}
		locations = append(locations, value)
	}
	return slices.Min(locations)
}

func extractSeeds(line string) []int {
	seeds := make([]int, 0)
	ids := strings.Fields(line[7:])
	for _, id := range ids {
		val := util.ConvertStringToInt(id)
		seeds = append(seeds, val)
	}
	return seeds
}

func extractAlmanac(lines []string) [][]Range {
	groups := make([][]Range, 0)
	ranges := make([]Range, 0)
	for i, line := range lines {
		if line == "" || strings.HasSuffix(line, ":") || len(lines)-1 == i {
			if len(ranges) > 0 {
				groups = append(groups, ranges)
				ranges = make([]Range, 0)
			}
			continue
		}
		ids := strings.Fields(line)
		r := Range{
			Dst:    util.ConvertStringToInt(ids[0]),
			Src:    util.ConvertStringToInt(ids[1]),
			Length: util.ConvertStringToInt(ids[2]),
		}
		ranges = append(ranges, r)

	}
	return groups
}
