package day4_part1

import (
	"strings"
)

func getAnswer(lines []string) int {
	ans := 0
	sb := strings.Builder{}
	for i, line := range lines {
		_, err := sb.WriteString(line)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(' ')
		if line == "" || i == len(lines)-1 {
			pass := sb.String()
			if isValid(pass) {
				ans++
			}
			sb.Reset()
		}
	}

	return ans
}

func isValid(s string) bool {
	//fmt.Printf("Checking: %s\n", s)
	parts := strings.Fields(s)
	m := make(map[string]bool)
	for _, part := range parts {
		field := part[0:3]
		m[field] = true
	}

	if len(m) == 8 {
		return true
	}

	_, ok := m["cid"]
	if !ok && len(m) == 7 {
		return true
	}

	return false
}
