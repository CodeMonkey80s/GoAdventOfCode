package day1_part2

import (
	"slices"

	"GoAdventOfCode/util"
)

func howManyTotalCalories(input []string) int {
	sum := 0
	total := 0
	calories := make([]int, 0)
	for i, line := range input {
		sum += util.ConvertStringToInt(line)
		if line == "" || len(input)-1 == i {
			calories = append(calories, sum)
			sum = 0
			continue
		}
	}

	slices.Sort(calories)
	total += calories[len(calories)-1]
	total += calories[len(calories)-2]
	total += calories[len(calories)-3]
	return total
}
