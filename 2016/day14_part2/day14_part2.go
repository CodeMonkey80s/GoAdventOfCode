package day14_part2

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strings"

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

func getHashStretched(prefix string, n int) string {
	hash := fmt.Sprintf("%s%d", prefix, n)
	for i := 0; i < 2017; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	return hash
}

func getAnswer(input string) int {

	hashes := make([]string, 0)
	for i := 0; i < 1000; i++ {
		hashes = append(hashes, getHashStretched(input, i))
	}

	keys := make([]int, 0)
	for i := 0; len(keys) < 64; i++ {

		currentHash := hashes[0]
		hashes = append(hashes, getHashStretched(input, i+1000))
		hashes = hashes[1:]

		ch, ok := hasThreeConsecutiveChars(currentHash)
		if ok {
			//fmt.Printf("Checking: %d %c\n", i, ch)
			pattern := regexp.MustCompile(fmt.Sprintf("[%s]{5}", string(ch)))
			if pattern.MatchString(strings.Join(hashes, ",")) {
				keys = append(keys, i)
				//fmt.Printf("Found: %d Key: %d\n", i, len(keys))
			}
		}
	}

	ans := keys[len(keys)-1]
	//fmt.Printf("Keys: %v\n", keys)
	fmt.Printf("Output: %v\n", ans)
	return ans
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
