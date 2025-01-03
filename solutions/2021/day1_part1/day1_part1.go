package day1_part1

import (
	"GoAdventOfCode/util"
)

func howManyMeasurementsLarger(lines []string) int {
	sum := 0
	last := 0
	for i := 0; i < len(lines); i++ {
		val := util.ConvertStringToInt(lines[i])
		if i > 0 {
			last = util.ConvertStringToInt(lines[i-1])
			if val > last {
				sum++
			}
		}
	}

	return sum
}
