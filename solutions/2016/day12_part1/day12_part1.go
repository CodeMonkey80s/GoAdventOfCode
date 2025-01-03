package day12_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string, registers map[string]int) int {

	n := 0
	for n = 0; n < len(lines); n++ {
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
			if util.IsNumber(x) {
				jump = true
			} else if registers[x] != 0 {
				jump = true
			}
			if jump {
				n += util.ConvertStringToInt(parts[2]) - 1
			}
		}
	}

	return registers["a"]
}
