package day1_part1

import (
	"GoAdventOfCode/util"
)

func getAnswerForPart1(lines []string) int {
	dial := 50
	var counter int
	for _, line := range lines {
		dir := line[0]
		numStr := line[1:]
		num := util.ConvertStringToInt(numStr)

		switch {
		case dir == 'L':
			dial = (dial - num + 100) % 100
		case dir == 'R':
			dial = (dial + num) % 100
		}
		if dial == 0 {
			counter++
		}
	}

	return counter
}
