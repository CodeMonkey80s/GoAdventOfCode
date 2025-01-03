package day11_part1

import (
	"fmt"
	"strings"
)

func getAnswer(input string) string {
	password := input
	i := 0
	for {
		password = incrementPassword(password)
		ok := isPasswordValid(password)
		if ok {
			fmt.Printf("Output: %d = %q\n", i, password)
			return password
		}
		i++
	}
	return ""
}

func incrementPassword(password string) string {

	runes := []rune(password)
	i := len(runes) - 1
	for {
		if runes[i] == 'z' {
			runes[i] = 'a'
			i--
			continue
		}
		runes[i]++
		break
	}

	return string(runes)
}

func isPasswordValid(password string) bool {

	ok := false
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i+1] == password[i+2]-1 {
			ok = true
			break
		}
	}
	if !ok {
		return false
	}

	if strings.ContainsAny(password, "iol") {
		return false
	}

	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i += 1
		}
	}
	if pairs < 2 {
		return false
	}

	return true
}
