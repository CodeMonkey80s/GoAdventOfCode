package day1_part1

import (
	"GoAdventOfCode/util"
)

func howManyTotalCalories(input []string) int {
	total := 0
	sum := 0
	for _, line := range input {
		if line == "" {
			total = max(total, sum)
			sum = 0
			continue
		}
		val := util.ConvertStringToInt(line)
		sum += val
	}

	return total
}
