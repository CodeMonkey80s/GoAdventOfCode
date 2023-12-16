package day2_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getTotal(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "x")
		a := util.ConvertStringToInt(parts[0])
		b := util.ConvertStringToInt(parts[1])
		c := util.ConvertStringToInt(parts[2])
		m := 1<<32 - 1
		m = min(m, a*b)
		m = min(m, b*c)
		m = min(m, c*a)
		t := (2*a*b + 2*b*c + 2*a*c) + m
		total += t
		//fmt.Printf("line: %q, %d %d %d, %d t: %d, total: %d\n", line, a, b, c, m, t, total)
	}
	return total
}
