package day3_part1

import (
	"strconv"
	"strings"
)

func sumOfAllOfThePartNumbers(schematic []string) int {
	sum := 0

	isDigit := func(ch byte) bool {
		return ch >= byte('0') && ch <= byte('9')
	}

	isSymbol := func(ch string) bool {
		if strings.ContainsAny(ch, "#$%&*+-/=@") {
			return true
		}
		return false
	}

	symbols := make(map[string]int)
	isPartNumber := func(sx int, sy int, size int) bool {
		for y := sy - 1; y <= sy+1; y++ {
			for x := sx - 1; x <= sx+size; x++ {
				if x >= 0 && y >= 0 && x < len(schematic[0]) && y < len(schematic) {
					ch := schematic[y][x]
					if ch == '.' {
						continue
					}
					if isDigit(ch) {
						continue
					}
					if isSymbol(string(ch)) {
						return true
					}
				}
			}
		}
		return false
	}

	for y := 0; y < len(schematic); y++ {
		process := false
		number := ""
		for x := 0; x < len(schematic[y]); x++ {
			ch := schematic[y][x]
			symbols[string(ch)]++
			if isDigit(ch) {
				number += string(ch)
			} else {
				process = true
			}
			if x == len(schematic)-1 {
				process = true
			}

			if process {
				size := len(number)
				if size > 0 {
					ok := isPartNumber(x-size, y, size)
					if ok {
						value, err := strconv.Atoi(number)
						if err != nil {
							panic(err)
						}
						sum += value
					}
				}
				number = ""
				process = false
			}
		}
	}

	return sum
}
