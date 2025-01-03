package day2_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(input string) int {

	codes := make([]int, 0)
	parts := strings.Split(input, ",")
	for _, number := range parts {
		value := util.ConvertStringToInt(number)
		codes = append(codes, value)
	}

	i := 0
loop:
	for {
		switch codes[i] {
		case 1:
			a := codes[codes[i+1]]
			b := codes[codes[i+2]]
			c := codes[i+3]
			codes[c] = a + b
			i += 4
			continue loop
		case 2:
			a := codes[codes[i+1]]
			b := codes[codes[i+2]]
			c := codes[i+3]
			codes[c] = a * b
			i += 4
			continue loop
		case 99:
			break loop
		}
		i++
	}

	//fmt.Printf("Codes: %+v\n", codes)
	return codes[0]
}
