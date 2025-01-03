package day10_part2

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

type operation struct {
	op string
	v  int
}

func getAnswer(input []string, maxCycles int) (int, int) {

	signal := 0
	register := 1

	operations := loadOperations(input)

	crt := loadCRT()

	px := -1
	py := 0

	for _, op := range operations {

		px++
		if px > 39 {
			py++
			px = 0
		}

		// fmt.Printf("i: %d, py=%d px=%d, %d\n", i, py, px, register)

		if register-1 == px || register == px || register+1 == px {
			crt[py][px] = '#'
		}

		switch op.op {
		case "noop":
		case "addx":
			register += op.v
		}
	}

	// printCRT(crt)

	return signal, register
}

func loadOperations(input []string) []operation {
	var operations []operation
	for _, line := range input {
		parts := strings.Fields(line)
		code := parts[0]
		switch {
		case code == "addx":
			op1 := operation{
				op: "noop",
			}
			operations = append(operations, op1)
			op2 := operation{
				op: "addx",
				v:  util.ConvertStringToInt(parts[1]),
			}
			operations = append(operations, op2)
		case code == "noop":
			op := operation{
				op: "noop",
			}
			operations = append(operations, op)
		}
	}

	return operations
}

func loadCRT() [][]rune {
	crt := make([][]rune, 6)
	for y := 0; y < 6; y++ {
		crt[y] = make([]rune, 40)
		for x := 0; x < 40; x++ {
			crt[y][x] = '.'
		}
	}

	return crt
}

func printCRT(crt [][]rune) {
	for y := 0; y < len(crt); y++ {
		for x := 0; x < len(crt[y]); x++ {
			fmt.Printf("%c", crt[y][x])
		}
		fmt.Println()
	}
}

/*
func getAnswer(input []string) int {

	crt := make([][]rune, 6)
	for y := 0; y < 6; y++ {
		crt[y] = make([]rune, 40)
		for x := 0; x < 40; x++ {
			crt[y][x] = '.'
		}
	}
	// displayCRT(crt)

	x := 1
	signal := 0
	cycle := 0

	a := false
	t := 0
	v := 0

	idx := 0

loop:
	for {

		cycle++

		if a && t == 0 {
			a = false
			x += v
			v = 0
			t = 0
		}

		fmt.Printf("cycle: %d, x=%d\n", cycle, x)
		if a && t > 0 {
			t--
		}

		// px := cycle%41 - 1
		// py := cycle / 41
		// if px < 0 {
		// 	px = 0
		// }
		// if py >= 0 && py < 6 {
		// 	fmt.Printf("cycle: %d, x=%d y=%d, %d\n", cycle, py, px, x)
		// 	if x == cycle-1 || x == cycle || x == cycle+1 {
		// 		crt[py][px] = '#'
		// 	} else {
		// 		crt[py][px] = '.'
		// 	}
		// }

		if a == true {
			continue loop
		}

		operation := getOperation(input, idx)
		parts := strings.Fields(operation)
		if operation == "" {
			break loop
		}
		opCode := parts[0]
		switch {
		case opCode == "addx" && a == false:
			a = true
			t = 1
			v = util.ConvertStringToInt(parts[1])
			idx++
		case opCode == "noop":
			idx++
		}

	}

	displayCRT(crt)

	return signal
}

*/
