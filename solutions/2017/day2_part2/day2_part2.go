package day2_part2

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	checksum := 0
	for _, line := range lines {
		numbers := strings.Fields(line)
		for _, number1 := range numbers {
			val1 := util.ConvertStringToInt(number1)
			for _, number2 := range numbers {
				val2 := util.ConvertStringToInt(number2)
				if val1%val2 == 0 && val1/val2 != 1 {
					checksum += val1 / val2
				}
			}
		}
	}

	fmt.Printf("checksum: %d\n", checksum)
	return checksum
}
