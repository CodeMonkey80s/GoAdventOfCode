package day1_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	for _, line1 := range lines {
		v1 := util.ConvertStringToInt(line1)
		for _, line2 := range lines {
			v2 := util.ConvertStringToInt(line2)
			if v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	return -1
}
