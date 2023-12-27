package day4_part1

import (
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

	// Must be six characters long
	if len(s) != 6 {
		//fmt.Printf("### Incorrect length\n")
		return false
	}

	// Two adjacent digits are the same (22 in 122345)
	found := false
	a := 0
	b := 1
	for i := 0; i < len(s)-1; i++ {
		if s[a] == s[b] {
			found = true
		}
		a++
		b++
	}
	if !found {
		//fmt.Printf("### No adjacent digits\n")
		return false
	}

	// Going from left to right, the digits never decrease
	a = 0
	b = 1
	for i := 0; i < len(s)-1; i++ {
		v1 := s[a] - '0'
		v2 := s[b] - '0'
		if v2 < v1 {
			//fmt.Printf("### Digits decrease\n")
			return false
		}
		a++
		b++
	}

	return true
}
