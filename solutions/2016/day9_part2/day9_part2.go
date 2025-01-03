package day9_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	total := 0
	for _, line := range lines {
		total += decompressString(line)
	}
	return total
}

func decompressString(s string) int {
	totalLength := 0

	if strings.Index(s, "(") == -1 {
		return len(s)
	}

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			idx := strings.Index(s[i:], ")")
			parts := strings.Split(s[i+1:i+idx], "x")
			length := util.ConvertStringToInt(parts[0])
			repeat := util.ConvertStringToInt(parts[1])
			segment := s[i+idx+1 : i+idx+1+length]
			if strings.Index(segment, "(") != -1 {
				totalLength += repeat * decompressString(segment)
			} else {
				totalLength += length * repeat
			}
			i += idx + length

		} else {
			totalLength++
		}

	}

	return totalLength
}
