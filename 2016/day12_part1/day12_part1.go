package day12_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	}

	n := 0
loop:
	for {
		line := lines[n]
		parts := strings.Fields(line)
		cmd := parts[0]
		x := parts[1]
		switch cmd {
		case "cpy":
			y := parts[2]
			if util.IsNumber(x) {
				registers[y] = util.ConvertStringToInt(x)
			} else {
				registers[y] = registers[x]
			}
		case "inc":
			registers[x]++
		case "dec":
			registers[x]--
		case "jnz":
			jump := false
			if x >= string('1') && x <= string('9') {
				jump = true
			} else if registers[x] != 0 {
				jump = true
			}
			if jump {
				n += util.ConvertStringToInt(parts[2])
				if n >= len(lines) {
					break
				}
				continue loop
			}
		}
		n++
		if n >= len(lines) {
			break
		}

	}

	fmt.Printf("Registers: %+v\n", registers)
	return registers["a"]
}
