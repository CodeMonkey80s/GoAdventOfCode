package day2_part2

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

	for y := 0; y <= 99; y++ {
		for x := 0; x <= 99; x++ {
			c := make([]int, len(codes))
			copy(c, codes)
			c[1] = y
			c[2] = x
			val := calculateValue(c)
			if val == 19690720 {
				return 100*y + x
			}
		}
	}
	//fmt.Printf("Codes: %+v\n", codes)
	return -1
}

func calculateValue(codes []int) int {
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
	return codes[0]
}
