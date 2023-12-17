package day1_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func sumOfAllTheCalibrationValues(lines []string) int {
	sum := 0
	nums := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, line := range lines {

		num1 := ""
		num2 := ""

		for i := 0; i < len(line); i++ {
			ch := line[i]
			if num1 == "" && ch >= '0' && ch <= '9' {
				num1 = string(ch)
				break
			}
			if num1 == "" {
				for j := 0; j < len(nums); j++ {
					if strings.HasPrefix(line[i:], nums[j]) {
						num1 = string('0' + byte(j) + 1)
						break
					}
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if num2 == "" && ch >= '0' && ch <= '9' {
				num2 = string(ch)
				break
			}
			if num2 == "" {
				for j := 0; j < len(nums); j++ {
					if strings.HasSuffix(line[:i+1], nums[j]) {
						num2 = string('0' + byte(j) + 1)
						break
					}
				}
			}
		}

		if num1 != "" && num2 != "" {
			number := num1 + num2
			val := util.ConvertStringToInt(number)
			sum += val
		}
	}

	return sum
}
