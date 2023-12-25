package day5_part1

import (
	"crypto/md5"
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getPassword(id string) string {

	password := ""

	n := 0
	for {
		word := id + util.ConvertIntToString(n)

		hash := md5.Sum([]byte(word))
		s := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(s, "00000") {
			password += string(s[5])
			fmt.Println(password)
			if len(password) == 8 {
				break
			}
		}
		n++
	}

	return password
}
