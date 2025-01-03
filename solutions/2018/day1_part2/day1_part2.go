package day1_part2

import (
	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	freq := map[int]int{
		0: 1,
	}
	f := 0
	for {
		for _, line := range lines {
			sign := line[0]
			val := util.ConvertStringToInt(line[1:])
			switch sign {
			case '+':
				f += val
			case '-':
				f -= val
			}
			freq[f]++
			if freq[f] == 2 {
				return f
			}
		}
	}
	return -1
}
