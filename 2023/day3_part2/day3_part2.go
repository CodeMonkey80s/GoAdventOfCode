package day3_part2

import (
	"strconv"
	"strings"
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
						value, err := strconv.Atoi(number)
						if err != nil {
							panic(err)
						}
						value2, ok := parts[pos]
						if ok {
							sum += (value * value2)
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

//func sumOfAllOfTheGearRatios(schematic []string) int {
//	sum := 0
//	for y := 0; y < len(schematic); y++ {
//		for x := 0; x < len(schematic[y]); x++ {
//			ch := schematic[y][x]
//			if ch == '*' {
//				numbers := make([]string, 0)
//				numbers = getNumbersFromBothSides(x, y, schematic, numbers)
//				numbers = getNumbers(x, y, schematic, numbers)
//				if len(numbers) == 2 {
//					//fmt.Printf("Numbers: %v\n", numbers)
//					value1, err := strconv.Atoi(numbers[0])
//					if err != nil {
//						panic(err)
//					}
//					value2, err := strconv.Atoi(numbers[1])
//					if err != nil {
//						panic(err)
//					}
//					ratio := value1 * value2
//					sum += ratio
//				}
//			}
//		}
//	}
//
//	return sum
//}

//func isDigit(ch byte) bool {
//	return ch >= byte('0') && ch <= byte('9')
//}

//func getNumbers(x int, y int, schematic []string, numbers []string) []string {
//
//	isStarAroundNumberBelow := func(sx int, sy int, size int) bool {
//		for y := sy; y <= sy+1; y++ {
//			for x := sx - size; x <= sx; x++ {
//				if x >= 0 && y >= 0 && x < len(schematic[0]) && y < len(schematic) {
//					ch := schematic[y][x]
//					if ch == '*' {
//						return true
//					}
//				}
//			}
//		}
//		return false
//	}

//isStarAroundNumberAbove := func(sx int, sy int, size int) bool {
//	for y := sy; y >= sy-1; y-- {
//		for x := sx - 1; x <= sx+size; x++ {
//			if x >= 0 && y >= 0 && x < len(schematic[0]) && y < len(schematic) {
//				ch := schematic[y][x]
//				if ch == '*' {
//					return true
//				}
//			}
//		}
//	}
//	return false
//}
//number := ""
//add := false
//for i := -3; i <= 3; i++ {
//	ch := schematic[y-1][x+i]
//	if isDigit(ch) {
//		number += string(ch)
//		continue
//	} else {
//		if number != "" {
//			add = true
//		}
//	}
//	if add {
//		size := len(number)
//		fmt.Printf("x: %d, y: %d, %s\n", x+i-size, y-1, number)
//		ok := isStarAroundNumberBelow(x+i-size, y-1, size)
//		if ok {
//			numbers = append(numbers, number)
//			add = false
//		}
//	}
//	number = ""
//}
//
//fmt.Printf("Numbers (top): %v\n", numbers)
//number = ""
//for i := -3; i <= 3; i++ {
//	ch := schematic[y+1][x+i]
//	if isDigit(ch) {
//		number += string(ch)
//	}
//	if ch == '.' || i == 3 {
//		if number != "" {
//			add = true
//		}
//	}
//	if add {
//		size := len(number)
//		ok := isStarAroundNumberAbove(x+i, y, size)
//		if ok {
//			numbers = append(numbers, number)
//			add = false
//			number = ""
//		}
//	}
//
//}
//fmt.Printf("Numbers (bottom): %v\n", numbers)
//	return numbers
//}

//func getNumbersFromBothSides(x int, y int, schematic []string, numbers []string) []string {
//	number := ""
//	for i := 1; i <= 3; i++ {
//		ch := schematic[y][x+i]
//		if i == 1 && (ch == '.' || ch == '*') {
//			break
//		}
//		if isDigit(ch) {
//			number += string(ch)
//			continue
//		}
//	}
//	if number != "" {
//		numbers = append(numbers, number)
//	}
//	number = ""
//	for i := -1; i >= -3; i-- {
//		ch := schematic[y][x+i]
//		if i == -1 && (ch == '.' || ch == '*') {
//			break
//		}
//		if isDigit(ch) {
//			number = string(ch) + number
//			continue
//		}
//	}
//	if number != "" {
//		numbers = append(numbers, number)
//	}
//	return numbers
//}
