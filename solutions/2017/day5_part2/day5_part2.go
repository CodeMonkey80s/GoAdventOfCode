package day5_part2

import (
	"GoAdventOfCode/util"
)

func getNumberOfSteps(lines []string) int {

	stack := make([]int, 0)
	for _, line := range lines {

		value := util.ConvertStringToInt(line)
		stack = append(stack, value)
	}

	steps := 0
	offset := 0
	for {
		steps++
		next := stack[offset]
		if stack[offset] >= 3 {
			stack[offset]--
		} else {
			stack[offset]++
		}
		offset += next
		if offset < 0 || offset >= len(stack) {
			break
		}
	}

	return steps
}
