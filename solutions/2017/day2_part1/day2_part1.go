package day2_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	checksum := 0
	for _, line := range lines {
		minVal := 1<<32 - 1
		maxVal := 0
		numbers := strings.Fields(line)
		for _, number := range numbers {
			val := util.ConvertStringToInt(number)
			minVal = min(minVal, val)
			maxVal = max(maxVal, val)
		}
		checksum += maxVal - minVal
	}

	return checksum
}
