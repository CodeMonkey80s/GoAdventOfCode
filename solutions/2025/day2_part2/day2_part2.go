package day2_part2

import (
	"GoAdventOfCode/util"
	"strings"
)

func getAnswerForPart2(lines []string) int {

	hasRepeatedSequence := func(s string) bool {

		parts := 2
		var sb strings.Builder
		for {
			for j := 0; j < parts; j++ {
				sb.WriteString(s[:len(s)/parts])
			}
			if sb.String() == s {
				return true
			}
			sb.Reset()
			parts++
			if parts > len(s) {
				break
			}
		}

		return false
	}

	ranges := strings.Split(lines[0], ",")

	var answer int
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		a := parts[0]
		b := parts[1]
		va := util.ConvertStringToInt(a)
		vb := util.ConvertStringToInt(b)

		for i := va; i <= vb; i++ {
			n := util.ConvertIntToString(i)
			ok := hasRepeatedSequence(n)
			if ok {
				answer += i
			}
		}
	}

	return answer
}
