package day4_part2

import (
	"strings"
)

func getAnswer(lines []string) int {
	total := 0
	for _, line := range lines {
		ok := isValidPassphrase(line)
		if ok {
			total++
		}
	}

	return total
}

func isValidPassphrase(line string) bool {
	parts := strings.Fields(line)
	for a := 0; a < len(parts); a++ {
		for b := 1; b < len(parts); b++ {
			if a != b {
				ok := isAnagrams(parts[a], parts[b])
				if ok {
					return false
				}
			}
		}
	}

	return true
}

func isAnagrams(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make(map[rune]int)
	for _, char := range word1 {
		freq1[char]++
	}

	freq2 := make(map[rune]int)
	for _, char := range word2 {
		freq2[char]++
	}

	for k, v := range freq1 {
		if freq2[k] != v {
			return false
		}
	}

	return true
}
