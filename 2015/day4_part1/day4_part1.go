package day4_part1

import (
	"crypto/md5"
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(secretKey string) string {
	i := 0
	for {
		number := util.ConvertIntToString(i)
		input := secretKey + number
		hash := md5.Sum([]byte(input))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), "00000") {
			return number
		}
		i++
		if i > 1_200_000 {
			break
		}
	}
	return ""
}
