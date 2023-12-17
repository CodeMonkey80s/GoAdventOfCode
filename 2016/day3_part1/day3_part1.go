package day3_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	total := 0
	for _, line := range lines {
		a, b, c := getSides(line)
		ok := isTriangle(a, b, c)
		if ok {
			total++
		}
	}

	return total
}

func getSides(line string) (int, int, int) {
	parts := strings.Fields(line)
	a := util.ConvertStringToInt(parts[0])
	b := util.ConvertStringToInt(parts[1])
	c := util.ConvertStringToInt(parts[2])

	return a, b, c
}

func isTriangle(a int, b int, c int) bool {
	if a+b > c && b+c > a && a+c > b {
		return true
	}

	return false
}
