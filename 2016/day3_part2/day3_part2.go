package day3_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	total := 0

	a := 0
	b := 0
	c := 0
	ok := false
	for i := 0; i < len(lines)-2; i += 3 {

		line1 := lines[i]
		line2 := lines[i+1]
		line3 := lines[i+2]

		parts1 := strings.Fields(line1)
		parts2 := strings.Fields(line2)
		parts3 := strings.Fields(line3)

		a = util.ConvertStringToInt(parts1[0])
		b = util.ConvertStringToInt(parts2[0])
		c = util.ConvertStringToInt(parts3[0])
		//fmt.Printf("%d %d %d\n", a, b, c)
		ok = isTriangle(a, b, c)
		if ok {
			total++
		}

		a = util.ConvertStringToInt(parts1[1])
		b = util.ConvertStringToInt(parts2[1])
		c = util.ConvertStringToInt(parts3[1])
		//fmt.Printf("%d %d %d\n", a, b, c)
		ok = isTriangle(a, b, c)
		if ok {
			total++
		}

		a = util.ConvertStringToInt(parts1[2])
		b = util.ConvertStringToInt(parts2[2])
		c = util.ConvertStringToInt(parts3[2])
		//fmt.Printf("%d %d %d\n", a, b, c)
		ok = isTriangle(a, b, c)
		if ok {
			total++
		}
	}

	return total
}

func isTriangle(a int, b int, c int) bool {
	if a+b > c && b+c > a && a+c > b {
		return true
	}

	return false
}
