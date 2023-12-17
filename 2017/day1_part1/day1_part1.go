package day1_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(input string) int {
	numbers := make([]int, 0)
	a := 0
	b := 1
	over := false
	for {
		ch1 := input[a]
		ch2 := input[b]
		if ch1 == ch2 {
			val := util.ConvertByteToInt(ch1)
			numbers = append(numbers, val)
		}
		if over {
			break
		}
		a++
		b++
		if b > len(input)-1 {
			a = len(input) - 1
			b = 0
			over = true
			continue
		}
	}
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
