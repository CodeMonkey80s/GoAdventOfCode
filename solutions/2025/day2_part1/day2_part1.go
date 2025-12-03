package day2_part1

import (
	"GoAdventOfCode/util"
	"fmt"
	"strings"
)

func getAnswerForPart1(lines []string) int {

	ranges := strings.Split(lines[0], ",")

	var answer int
	for _, r := range ranges {

		parts := strings.Split(r, "-")
		a := parts[0]
		b := parts[1]
		va := util.ConvertStringToInt(a)
		vb := util.ConvertStringToInt(b)

	loop:
		for i := va; i <= vb; i++ {
			n := util.ConvertIntToString(i)
			if len(n)%2 == 0 {
				for j := 0; j < len(n)/2; j++ {
					if n[j] != n[j+len(n)/2] {
						continue loop
					}
				}
				fmt.Printf("a: %v, b: %v, inc: %d\n", a, b, i)
				answer += i
			}
		}

	}

	return answer
}
