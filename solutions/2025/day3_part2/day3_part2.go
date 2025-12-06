package day3_part2

import (
	"GoAdventOfCode/util"
	"strings"
)

func getAnswerForPart2(lines []string) int {

	findMax := func(a int, b int, s string) (byte, int) {
		var (
			ch  byte
			idx int
		)
		for i := a; i < b; i++ {
			if s[i] > ch {
				ch = s[i]
				idx = i
			}
		}

		return ch, idx
	}

	var (
		joltage int
		sb      strings.Builder
	)

	const batteries = 12

	for _, s := range lines {

		var idx int
		var ch byte
		for i := 0; i < batteries; i++ {

			a := idx
			b := len(s) - batteries + i + 1
			ch, idx = findMax(a, b, s)
			sb.WriteByte(ch)
			idx++

		}

		joltage += util.ConvertStringToInt(sb.String())
		sb.Reset()
	}

	return joltage
}
