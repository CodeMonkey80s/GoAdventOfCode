package day1_part2

import (
	"GoAdventOfCode/util"
)

func howManyMeasurementsLarger(lines []string) int {
	sum := 0
	for i := 0; i < len(lines)-3; i++ {
		sa := getWindowSum(i+0, lines)
		sb := getWindowSum(i+1, lines)
		//fmt.Printf("%v = %v %v\n", i, sa, sb)
		if sb > sa {
			sum++
		}
	}
	return sum
}

func getWindowSum(n int, lines []string) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += util.ConvertStringToInt(lines[n+i])
	}
	return sum
}
