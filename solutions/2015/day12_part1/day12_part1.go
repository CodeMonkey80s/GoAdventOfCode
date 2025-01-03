package day12_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(json string) int {

	sum := 0
	digits := make([]rune, 0)
	value := 0
	for _, char := range json {
		if isDigit(char) || char == '-' {
			digits = append(digits, char)
		} else {
			if len(digits) > 0 {
				value = util.ConvertStringToInt(string(digits))
				sum += value
				value = 0
				digits = digits[:0:0]
			}
		}

	}

	return sum
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
