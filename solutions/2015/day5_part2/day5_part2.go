package day5_part2

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
	ok1 := doesStringContainPairOfAnyTwoLetters(s)
	ok2 := doesStringContainLetterThatRepeats(s)
	//fmt.Printf("%t = %t\n", ok1, ok2)
	return ok1 && ok2
}

func doesStringContainPairOfAnyTwoLetters(s string) bool {
	n := 0
	for i := 0; i < len(s)-1; i++ {
		word := s[i : i+2]
		idx := strings.Index(s[i+2:], word)
		if idx != -1 {
			//fmt.Printf("ok1: %v\n", word)
			n += 2
		}
	}
	return n >= 2
}

func doesStringContainLetterThatRepeats(s string) bool {
	for i := 0; i < len(s); i++ {
		for w := 0; w < len(s)-i-1; w++ {
			word := s[i+w : i+w+1]
			idx := strings.Index(s[i+w+2:], word)
			if idx == 0 {
				return true
			}
		}
	}
	return false
}
