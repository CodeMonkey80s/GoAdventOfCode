package day2_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	ans := 0
	for _, line := range lines {
		lo, hi, letter, password := getPolicy(line)
		//fmt.Printf("lo: %v, hi: %v, letter: %q, password: %q\n", lo, hi, letter, password)
		ok := isPasswordValid(lo, hi, letter, password)
		if ok {
			ans++
		}
	}

	return ans
}

func getPolicy(line string) (int, int, rune, string) {
	letter := '-'
	password := ""
	lo, hi := 0, 0
	idx := strings.Index(line, ":")
	letter = rune(line[idx-1])
	numbers := line[:idx-2]
	password = line[idx+2:]
	parts := strings.Split(numbers, "-")
	lo = util.ConvertStringToInt(parts[0])
	hi = util.ConvertStringToInt(parts[1])
	return lo, hi, letter, password
}

func isPasswordValid(lo int, hi int, letter rune, password string) bool {
	n := 0
	if rune(password[lo-1]) == letter {
		n++
	}
	if rune(password[hi-1]) == letter {
		n++
	}
	if n == 1 {
		return true
	}

	return false
}
