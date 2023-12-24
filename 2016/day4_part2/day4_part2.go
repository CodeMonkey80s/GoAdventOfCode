package day4_part2

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	for _, line := range lines {
		id, password := decryptPassword(line)
		if strings.Contains(password, "north") {
			fmt.Printf("id: %d, p: %q\n", id, password)
			return id
		}
	}
	return 0
}

func decryptPassword(s string) (int, string) {
	l := len(s)
	id := util.ConvertStringToInt(s[l-10 : l-7])

	s = s[:l-10]
	s = strings.TrimSuffix(s, "-")
	s = strings.Replace(s, "-", " ", -1)

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	letters := []rune(s)
	p := rotateString(letters, id)

	return id, p
}

func rotateString(letters []rune, n int) string {
	for j := 0; j < n; j++ {
		for i := range letters {
			if letters[i] == ' ' {
				continue
			}
			if letters[i] == 'z' {
				letters[i] = 'a'
				continue
			}
			letters[i]++
		}
	}
	return string(letters)
}
