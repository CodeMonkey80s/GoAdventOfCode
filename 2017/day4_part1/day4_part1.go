package day4_part1

import "strings"

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
	freq := make(map[string]int)
	parts := strings.Fields(line)
	for _, part := range parts {
		_, ok := freq[part]
		if ok {
			return false
		}
		freq[part]++

	}

	return true
}
