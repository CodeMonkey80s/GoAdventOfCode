package day6_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func howManyLanternfish(input string) int {
	freq := make(map[int]int)
	parts := strings.Split(input, ",")
	for _, number := range parts {
		n := util.ConvertStringToInt(number)
		freq[n]++
	}
	day := 1
	after := 256
	for {
		newFreq := make(map[int]int)
		for k, n := range freq {
			if k == 0 {
				newFreq[6] += n
				newFreq[8] += n
			} else {
				newFreq[k-1] += n
			}
		}
		freq = newFreq
		day++
		if day > after {
			break
		}
	}
	total := 0
	for _, v := range freq {
		total += v
	}
	return total
}
