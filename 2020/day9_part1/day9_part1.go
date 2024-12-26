package day9_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(input []string, size int) int {
	n := size
	for {
		valid := false
		n3 := util.ConvertStringToInt(input[n])
		for i := n - size; i < n; i++ {
			n1 := util.ConvertStringToInt(input[i])
			for j := n - size; j < n; j++ {
				if i == j {
					continue
				}
				n2 := util.ConvertStringToInt(input[j])
				if n1+n2 == n3 {
					valid = true
				}
			}
		}

		if !valid {
			return n3
		}

		n++
		if n > len(input) {
			break
		}
	}

	return -1
}
