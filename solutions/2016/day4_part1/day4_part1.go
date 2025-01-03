package day4_part1

import (
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	ans := 0
	for _, line := range lines {
		id, ok := isValidPassword(line)
		if ok {
			val := util.ConvertStringToInt(id)
			ans += val
		}
	}

	return ans
}

func isValidPassword(s string) (string, bool) {
	l := len(s)
	id := s[l-10 : l-7]
	checksum := s[l-6 : l-1]
	//fmt.Printf("id: %q, checksum: %q\n", id, checksum)
	//fmt.Println("_")

	s = s[:l-10]
	s = strings.Replace(s, "-", "", -1)

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	//fmt.Printf("freq: %q\n", freq)
	letters := []rune(s)
	slices.SortFunc(letters, func(a, b rune) int {
		if freq[a] == freq[b] {
			if a < b {
				return -1
			}
			return 1
		}
		if freq[a] < freq[b] {
			return 1
		} else {
			return -1
		}
	})
	letters = slices.Compact(letters)

	if string(letters)[:5] == checksum {
		return id, true
	}

	//fmt.Printf("letters: %q\n", letters)

	return "", false
}
