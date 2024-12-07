package day7_part1

import (
	"fmt"
	"math"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	sum := 0
outer:
	for _, line := range lines {
		parts := strings.Fields(line)

		total := util.ConvertStringToInt(strings.TrimSuffix(parts[0], ":"))

		var numbers []int
		for _, num := range parts[1:] {
			numbers = append(numbers, util.ConvertStringToInt(num))
		}

		size := len(numbers) - 1
		combinations := getCombinations(size)
		for _, combo := range combinations {
			val := checkCombination(numbers, combo)
			if val == total {
				sum += total
				continue outer
				// fmt.Printf("sum: %s\n", sum.String())
				// fmt.Printf("max: %d\n", uint64(math.MaxUint64))
			}
		}

	}

	return sum
}

func getCombinations(size int) []string {
	var combinations []string

	v := int(math.Pow(2, float64(size))) - 1
	for i := 0; i <= v; i++ {
		s := fmt.Sprintf("%0*b", size, i)
		combinations = append(combinations, s)
	}

	return combinations
}

func checkCombination(numbers []int, combo string) int {
	sum := numbers[0]
	for i, num := range numbers[1:] {
		if combo[i] == '0' {
			sum += num
		} else {
			sum *= num
		}
	}

	return sum
}
