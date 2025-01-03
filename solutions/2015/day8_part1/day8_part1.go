package day8_part1

import (
	"strconv"
	"unicode/utf8"
)

func getAnswer(lines []string) int {
	ans := 0
	for _, line := range lines {
		c1, c2 := countBytesAndRunes(line)
		ans += c1 - c2
	}
	return ans
}

func countBytesAndRunes(characters string) (int, int) {
	bytesCount := len(characters)
	chars, err := strconv.Unquote(characters)
	if err != nil {
		panic(err)
	}
	runeCount := utf8.RuneCountInString(chars)
	return bytesCount, runeCount
}
