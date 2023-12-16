package day1_part1

import "GoAdventOfCode/util"

func getAnswer(lines []string) int {
	sum := 0
	for _, line := range lines {
		sign := line[0]
		val := util.ConvertStringToInt(line[1:])
		switch sign {
		case '+':
			sum += val
		case '-':
			sum -= val
		}
	}
	return sum
}
