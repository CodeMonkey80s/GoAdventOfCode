package day1_part2

import "GoAdventOfCode/util"

func getAnswer(lines []string) int {
	for _, line1 := range lines {
		v1 := util.ConvertStringToInt(line1)
		for _, line2 := range lines {
			v2 := util.ConvertStringToInt(line2)
			for _, line3 := range lines {
				v3 := util.ConvertStringToInt(line3)
				if v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	return -1
}
