package day1_part1

import (
	"unicode"

	"GoAdventOfCode/util"
)

func sumOfAllTheCalibrationValues(lines []string) int {
	sum := 0
	for _, line := range lines {
		digits := make([]rune, 0)
		for _, char := range line {
			if unicode.IsDigit(char) {
				digits = append(digits, char)
			}
		}
		switch {
		case len(digits) == 1:
			digits = append(digits, digits...)
		case len(digits) >= 3:
			digits = append([]rune{digits[0]}, digits[len(digits)-1])
		}
		sum += util.ConvertStringToInt(string(digits))
	}

	return sum
}
