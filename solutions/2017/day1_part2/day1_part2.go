package day1_part2

import (
	"GoAdventOfCode/util"
)

func getAnswer(input string) int {
	numbers := make([]int, 0)

	length := len(input)
	step := length / 2

	for i := 0; i < length; i++ {
		ch1 := input[i]
		idx := (i + step) % length
		ch2 := input[idx]
		if ch1 == ch2 {
			val := util.ConvertByteToInt(input[i])
			numbers = append(numbers, val)
		}
	}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
