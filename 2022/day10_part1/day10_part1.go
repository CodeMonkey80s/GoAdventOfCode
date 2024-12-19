package day10_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type op struct {
	t int
	v int
}

func getAnswer(input []string) int {

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

		if a && t > 0 {
			t--
		}

		if a && t == 0 {
			a = false
			x += v
			v = 0
			t = 0
		}

		switch {
		case cycle == 20:
			fallthrough
		case cycle == 60:
			fallthrough
		case cycle == 100:
			fallthrough
		case cycle == 140:
			fallthrough
		case cycle == 180:
			fallthrough
		case cycle == 220:
			// fmt.Printf("cycle: %d, x: %d\n", cycle, x)
			signal += cycle * x
		}

		if a == true {
			continue loop
		}

		operation := getOperation(input, idx)
		parts := strings.Fields(operation)
		// fmt.Printf("cycle: %d, op: %q\n", cycle, opCode)
		if operation == "" {
			break loop
		}
		opCode := parts[0]
		switch {
		case opCode == "addx" && a == false:
			a = true
			t = 2
			v = util.ConvertStringToInt(parts[1])
			idx++
		case opCode == "noop":
			idx++
		}

	}

	// fmt.Printf("x: %d\n", x)

	return signal
}

func getOperation(input []string, op int) string {
	if op >= len(input) {
		return ""
	}
	return input[op]
}
