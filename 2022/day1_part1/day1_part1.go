package day1_part1

import "strconv"

func howManyTotalCalories(input []string) int {
	total := 0
	sum := 0
	for _, line := range input {
		if line == "" {
			total = max(total, sum)
			sum = 0
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += val
	}

	return total
}
