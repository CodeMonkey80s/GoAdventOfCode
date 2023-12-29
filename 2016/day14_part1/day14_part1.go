package day14_part1

import (
	"crypto/md5"
	"fmt"

	"GoAdventOfCode/util"
)

func hasThreeConsecutiveChars(hash string) (rune, bool) {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return rune(hash[i]), true
		}
	}
	return 0, false
}

func hasFiveConsecutiveChars(hash string, ch rune) bool {
	for i := 0; i < len(hash)-5; i++ {
		if ch == rune(hash[i+1]) &&
			ch == rune(hash[i+2]) &&
			ch == rune(hash[i+3]) &&
			ch == rune(hash[i+4]) &&
			ch == rune(hash[i+5]) {
			return true
		}
	}
	return false
}

func getHash(prefix string, n int) string {
	s := prefix + util.ConvertIntToString(n)
	h := md5.Sum([]byte(s))
	hash := fmt.Sprintf("%x", h)
	return hash
}

func getAnswer(input string) int {
	n := 0
	list := make([]string, 0)
	for {
		hash := getHash(input, n)

		ch, ok := hasThreeConsecutiveChars(hash)
		if ok {
			//fmt.Printf("Checking: %q (%d-%d)\n", ch, n, n+1000)
			for j := 1; j <= 1000; j++ {
				m := n + j
				hash := getHash(input, m)
				ok2 := hasFiveConsecutiveChars(hash, ch)
				if ok2 {
					//fmt.Printf("Found: %d at %d %q\n", n, m, hash)
					list = append(list, hash)
					break
				}
			}
		}
		if len(list) == 64 {
			break
		}

		n++
	}
	fmt.Printf("Output: %v\n", n)

	return n
}

func firstHashWith(input string, f func(hash string) bool) int {
	n := 0
	for {
		number := util.ConvertIntToString(n)
		s := input + number
		h := md5.Sum([]byte(s))
		hash := fmt.Sprintf("%x", h)
		if f(hash) {
			return n
		}
		n++
	}
	return -1
}
