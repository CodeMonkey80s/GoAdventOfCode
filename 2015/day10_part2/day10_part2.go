package day10_part2

import (
	"bytes"

	"GoAdventOfCode/util"
)

const (
	numberOfIterations = 50
)

func getAnswer(input string) int {
	for i := 0; i < numberOfIterations; i++ {
		input = lookAndSay(input)
	}

	return len(input)
}

func lookAndSay(input string) string {
	buf := bytes.Buffer{}

	var char uint8
	var count int

	for i := range input {
		switch {
		case count == 0:
			char = input[i]
			count++
		case i > 0 && input[i] == input[i-1]:
			count++
		default:
			c := util.ConvertIntToString(count)
			buf.WriteString(c)
			buf.WriteString(string(char))
			char = input[i]
			count = 1
		}
	}

	c := util.ConvertIntToString(count)
	buf.WriteString(c)
	buf.WriteString(string(char))

	return buf.String()
}
