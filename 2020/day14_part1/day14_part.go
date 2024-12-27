package day14_part1

import (
	"fmt"
	"strconv"
	"strings"

	"GoAdventOfCode/util"
)

const (
	bitSize = 36
)

func getAnswer(input []string) int {

	memory := make(map[string]string)

	var mask string

	for _, line := range input {
		if strings.HasPrefix(line, "mask = ") {
			mask = strings.ReplaceAll(line, "mask = ", "")
			continue
		}

		s := strings.TrimPrefix(line, "mem[")
		s = strings.ReplaceAll(s, "]", "")
		parts := strings.Split(s, " = ")

		bank := parts[0]
		val := fmt.Sprintf("%b", util.ConvertStringToInt(parts[1]))

		value := make([]rune, bitSize)
		a := 0
		for i := 0; i < bitSize; i++ {
			value[i] = '0'
			if i > bitSize-len(val)-1 {
				value[i] = rune(val[a])
				a++
			}
		}

		result := make([]rune, bitSize)
		for i, ch := range mask {
			switch {
			case ch == 'X':
				result[i] = rune(value[i])
			case ch == '0':
				result[i] = '0'
			case ch == '1':
				result[i] = '1'
			}
		}

		memory[bank] = string(result[:])

		// fmt.Printf("bank: %s, v: %s\n", bank, string(value))
		// fmt.Printf("bank: %s, m: %s\n", bank, mask)
		// fmt.Printf("bank: %s, r: %s\n", bank, string(result))
	}

	sum := 0
	for _, value := range memory {
		v, _ := strconv.ParseInt(value, 2, 64)
		sum += int(v)
		// fmt.Printf("bank: %s, %s, %d\n", bank, value, v)
	}

	return sum
}
