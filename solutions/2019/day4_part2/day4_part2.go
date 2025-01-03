package day4_part2

import (
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(input string) int {

	ans := 0

	parts := strings.Split(input, "-")
	lower := util.ConvertStringToInt(parts[0])
	upper := util.ConvertStringToInt(parts[1])

	for i := lower; i <= upper; i++ {
		ok := verify(i)
		if ok {
			ans++
		}
	}

	return ans
}

func verify(n int) bool {

	s := util.ConvertIntToString(n)
	sl := []rune(s[:])
	slices.Sort(sl)

	if s != string(sl) {
		return false
	}

	freq := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		freq[ch]++
	}
	for _, v := range freq {
		if v == 2 {
			return true
		}
	}

	return false
}
