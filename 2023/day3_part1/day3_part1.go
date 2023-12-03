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

	/*
			map[#:1 $:1 *:3 +:1 .:66 1:3 2:1 3:3 4:3 5:6 6:5 7:3 8:2 9:2]
		    map[#:38 $:54 %:38 &:56 *:390 +:34 -:47 .:15297 /:41 0:235 1:354 2:387 3:374 4:371 5:356 6:361 7:386 8:365 9:347 =:33 @:36]
	*/
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
					//fmt.Printf("(%d, %d) = %q\n", x, y, ch)
					if ch == '.' {
						continue
					}
					if isDigit(ch) {
						continue
					}
					if isSymbol(string(ch)) {
						//fmt.Printf("%c\n", ch)
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
					//fmt.Printf("number: %s, %d, %t\n", number, size, ok)
					if ok {
						value, err := strconv.Atoi(number)
						if err != nil {
							panic(err)
						}
						//fmt.Printf("v: %s = %d\n", number, value)
						sum += value
					}
				}
				number = ""
				process = false
			}
		}
	}

	//fmt.Printf("Symbols: %v\n", symbols)
	return sum
}
