package day1

import (
	"strconv"
	"strings"
)

func getAnswerForPart2(lines []string) int {

	numbers1 := make([]int, 0, len(lines))
	numbers2 := make([]int, 0, len(lines))

	for _, line := range lines {

		parts := strings.Fields(line)
		num1 := parts[0]
		num2 := parts[1]

		v1, _ := strconv.Atoi(num1)
		v2, _ := strconv.Atoi(num2)

		numbers1 = append(numbers1, v1)
		numbers2 = append(numbers2, v2)
	}

	getScore := func(n int) int {
		score := 0
		for _, v := range numbers2 {
			if n == v {
				score++
			}
		}
		return score
	}

	sum := 0
	for i := range numbers1 {
		sum += numbers1[i] * getScore(numbers1[i])
	}

	return sum
}
