package day5_part2

import (
	"crypto/md5"
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getPassword(id string) string {
	password := [8]rune{}
	n := 0
	c := 0
	for {
		word := id + util.ConvertIntToString(n)
		hash := md5.Sum([]byte(word))
		s := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(s, "00000") {
			position := s[5] - '0'
			if position > 7 {
				n++
				continue
			}
			if password[position] != 0 {
				n++
				continue
			}
			char := s[6]
			password[position] = rune(char)
			c++
			if c == 8 {
				break
			}
		}
		n++
	}

	return string(password[:])
}
