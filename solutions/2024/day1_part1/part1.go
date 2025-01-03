package day1_part1

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func getAnswerForPart1(lines []string) int {

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

	slices.Sort(numbers1)
	slices.Sort(numbers2)

	sum := 0
	for i := range numbers1 {
		diff := math.Abs(float64(numbers2[i] - numbers1[i]))
		sum += int(diff)
	}

	return sum
}
