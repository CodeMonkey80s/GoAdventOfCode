package day8_part2

import (
	"strconv"
	"unicode/utf8"
)

func getAnswer(lines []string) int {
	ans := 0
	for _, line := range lines {
		c1, _ := countBytesAndRunes(line)
		encoded := strconv.Quote(line)
		c3, _ := countBytesAndRunes(encoded)
		ans += c3 - c1
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
