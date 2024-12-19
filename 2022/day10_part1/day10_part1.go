package day10_part1

import (
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

	for i, op := range operations {
		cycle := i + 1

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
			s := cycle * register
			signal += s
			// fmt.Printf("period: %d, %d, signal: %d\n", cycle, register, s)
		}

		switch op.op {
		case "noop":
		case "addx":
			register += op.v
		}

		// fmt.Printf("cycle: %d, op: %v, register: %d\n", cycle, op, register)

	}

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
