package day5_part1

import (
	"strings"
)

func getAnswer(input []string) int {
	total := 0
	for _, s := range input {
		if ok := isStringNice(s); ok {
			total++
		}
	}
	return total
}

func isStringNice(s string) bool {
	ok1 := doesStringContainThreeVowels(s)
	ok2 := doesStringContainDoubleLetter(s)
	ok3 := doesStringContainBadWord(s)
	//fmt.Printf("%t, %t, %t\n", ok1, ok2, ok3)
	return ok1 && ok2 && !ok3
}

func doesStringContainThreeVowels(s string) bool {
	freq := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
	for _, char := range s {
		_, ok := freq[char]
		if ok {
			freq[char]++
		}
	}
	n := 0
	for _, v := range freq {
		n += v
	}
	return n >= 3
}

func doesStringContainDoubleLetter(s string) bool {
	a := 0
	b := 1
	n := 0
	for i := 0; i < len(s)-1; i++ {
		if s[a] == s[b] {
			n++
		}
		a++
		b++
	}
	return n >= 1
}

func doesStringContainBadWord(s string) bool {
	words := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, word := range words {
		idx := strings.Index(s, word)
		if idx != -1 {
			return true
		}
	}
	return false
}
