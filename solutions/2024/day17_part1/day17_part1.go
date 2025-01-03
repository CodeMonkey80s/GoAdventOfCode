package day17_part1

import (
	"math"
	"strings"

	"GoAdventOfCode/util"
)

func getOutput(lines []string) (map[rune]int, string) {

	registers := map[rune]int{
		'a': util.ConvertStringToInt(strings.Replace(lines[0], "Register A: ", "", -1)),
		'b': util.ConvertStringToInt(strings.Replace(lines[1], "Register B: ", "", -1)),
		'c': util.ConvertStringToInt(strings.Replace(lines[2], "Register C: ", "", -1)),
	}

	input := strings.Split(strings.Replace(lines[4], "Program: ", "", -1), ",")

	var program []int
	for _, n := range input {
		v := util.ConvertStringToInt(n)
		program = append(program, v)
	}

	var numbers []int

	combo := func(v int) int {
		switch v {
		case 4:
			return registers['a']
		case 5:
			return registers['b']
		case 6:
			return registers['c']
		default:
			return v
		}
	}

	ip := 0
loop:
	for ip < len(program) {
		opcode := program[ip]

		switch opcode {
		case 0:
			out := int(float64(registers['a']) / math.Pow(2, float64(combo(program[ip+1]))))
			registers['a'] = out
		case 1:
			out := registers['b'] ^ program[ip+1]
			registers['b'] = out
		case 2:
			out := combo(program[ip+1]) % 8
			registers['b'] = out
		case 3:
			if registers['a'] != 0 {
				ip = program[ip+1]
				continue loop
			}
		case 4:
			out := registers['b'] ^ registers['c']
			registers['b'] = out
		case 5:
			out := combo(program[ip+1]) % 8
			numbers = append(numbers, out)
		case 6:
			out := int(float64(registers['a']) / math.Pow(2, float64(combo(program[ip+1]))))
			registers['b'] = out
		case 7:
			out := int(float64(registers['a']) / math.Pow(2, float64(combo(program[ip+1]))))
			registers['c'] = out
		}

		ip += 2
	}

	var output []string
	for _, n := range numbers {
		s := util.ConvertIntToString(n)
		output = append(output, s)
	}

	return registers, strings.Join(output, ",")
}
