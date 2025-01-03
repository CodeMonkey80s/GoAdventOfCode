package day15_part2

import (
	"regexp"

	"GoAdventOfCode/util"
)

type disc struct {
	pos    int
	maxPos int
}

func getAnswer(input []string) int {

	discs := loadDiscs(input)
	for turn := 0; ; turn++ {

		isDiscPassed := true
		for i := 0; i < len(discs); i++ {
			hole := (discs[i].pos + turn + i) % discs[i].maxPos
			if hole != 0 {
				isDiscPassed = false
				break
			}
		}

		if isDiscPassed {
			return turn - 1
		}
	}
}

func loadDiscs(input []string) []disc {

	var discs []disc

	var (
		reDiscPositions = regexp.MustCompile(`has ([0-9]+) positions`)
		reDiscPosition  = regexp.MustCompile(`position ([0-9]+)\.`)
	)

	for _, line := range input {
		d := disc{}
		if matches := reDiscPositions.FindStringSubmatch(line); len(matches) > 0 {
			d.maxPos = util.ConvertStringToInt(matches[1])
		}
		if matches := reDiscPosition.FindStringSubmatch(line); len(matches) > 0 {
			d.pos = util.ConvertStringToInt(matches[1])
		}

		discs = append(discs, d)
	}

	return discs
}
