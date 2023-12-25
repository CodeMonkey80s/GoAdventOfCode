package day9_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	total := 0
	for _, line := range lines {
		s := decompressString(line)
		total += len(s)

	}
	return total
}

func decompressString(s string) string {

	output := ""
	for i := 0; i < len(s); i++ {

		char := s[i]
		if char == '(' {
			idx := strings.Index(s[i:], ")")
			parts := strings.Split(s[i+1:i+idx], "x")
			length := util.ConvertStringToInt(parts[0])
			repeat := util.ConvertStringToInt(parts[1])

			segment := s[i+idx+1 : i+idx+1+length]

			//fmt.Printf("length: %v\n", length)
			//fmt.Printf("repeat: %v\n", repeat)
			//fmt.Printf("parts: %v\n", parts)
			//fmt.Printf("segment: %v\n", segment)

			for j := 0; j < repeat; j++ {
				output += segment
			}
			i += idx + length
		} else {
			output += string(char)
		}

	}

	return output
}
