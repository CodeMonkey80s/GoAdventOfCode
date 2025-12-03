package day1_part2

import (
	"GoAdventOfCode/util"
)

func getAnswerForPart2(lines []string) int {
	dial := 50
	var counter int
	for _, line := range lines {
		dir := line[0]
		numStr := line[1:]
		num := util.ConvertStringToInt(numStr)

		switch {
		case dir == 'L':
			for i := 0; i < num; i++ {
				dial--
				if dial == 0 {
					counter++
				}
				if dial < 0 {
					dial = 99
				}
			}
		case dir == 'R':
			for i := 0; i < num; i++ {
				dial++
				if dial > 99 {
					dial = 0
				}
				if dial == 0 {
					counter++
				}
			}
		}
	}

	return counter
}
