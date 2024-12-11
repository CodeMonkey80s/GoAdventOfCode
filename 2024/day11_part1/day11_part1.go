package day11_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(stones []int, n int) ([]int, int) {
	for i := 0; i < n; i++ {
		stones = blink(stones)
	}

	return stones, len(stones)
}

func blink(stonesOrg []int) []int {
	var stones []int
	for _, stone := range stonesOrg {

		if stone == 0 {
			stones = append(stones, 1)
			continue
		}

		digits, ok := hasEvenDigits(stone)
		if ok {
			num1 := digits[0 : len(digits)/2]
			num2 := digits[len(digits)/2:]
			v1 := util.ConvertStringToInt(num1)
			v2 := util.ConvertStringToInt(num2)
			stones = append(stones, v1)
			stones = append(stones, v2)
			continue
		}

		stones = append(stones, stone*2024)
	}

	return stones
}

func hasEvenDigits(stone int) (string, bool) {
	digits := util.ConvertIntToString(stone)
	if len(digits)%2 == 0 {
		return digits, true
	}

	return "", false
}
