package day6_part2

import (
	"GoAdventOfCode/util"
	"strings"
)

func getAnswerForPart2(lines []string) int {

	symbols := lines[len(lines)-1]
	a := 0
	b := 0

	var total int

	for {
		for i := 1; i < 6; i++ {
			idx := a + i
			if idx >= len(lines[0]) {
				b = idx
				break
			}
			if idx < len(lines[0]) && (symbols[idx] == '+' || symbols[idx] == '*') {
				b = idx - 1
				break
			}
		}

		var colSum int
		symbol := symbols[a]

		for x := a; x < b; x++ {
			var sb strings.Builder
			for y := 0; y < len(lines)-1; y++ {
				ch := lines[y][x]
				if ch >= '0' && ch <= '9' {
					sb.WriteByte(lines[y][x])
				}
			}

			v := util.ConvertStringToInt(sb.String())
			sb.Reset()

			switch {
			case colSum == 0:
				colSum = v
			case symbol == '+':
				colSum += v
			case symbol == '*':
				colSum *= v
			}
		}

		total += colSum

		a = b + 1
		if a >= len(lines[0]) {
			break
		}
	}

	return total
}
