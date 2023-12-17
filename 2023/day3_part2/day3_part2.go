package day3_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func sumOfAllOfTheGearRatios(schematic []string) int {
	sum := 0

	isDigit := func(ch byte) bool {
		return ch >= byte('0') && ch <= byte('9')
	}

	isAsterisk := func(ch string) bool {
		return strings.ContainsAny(ch, "*")
	}

	isSymbol := func(ch string) bool {
		return strings.ContainsAny(ch, "#$%&+-/=@")
	}

	type position struct {
		X int
		Y int
	}

	parts := make(map[position]int)

	symbols := make(map[string]int)
	isPartNumber := func(sx int, sy int, size int) (bool, position) {
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
					if isAsterisk(string(ch)) {
						return true, position{X: x, Y: y}
					}
					if isSymbol(string(ch)) {
						return true, position{X: x, Y: y}
					}
				}
			}
		}
		return false, position{X: 0, Y: 0}
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
					isPart, pos := isPartNumber(x-size, y, size)
					if isPart {
						value := util.ConvertStringToInt(number)
						value2, ok := parts[pos]
						if ok {
							sum += value * value2
						} else {
							parts[pos] = value
						}
					}
				}
				number = ""
				process = false
			}
		}
	}

	return sum
}
