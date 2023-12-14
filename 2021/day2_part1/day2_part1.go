package day2_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func horizontalPosition(commands []string) int {
	pos := 0
	depth := 0
	for _, command := range commands {
		parts := strings.Fields(command)
		cmd := parts[0]
		val := util.ConvertStringToInt(parts[1])
		switch cmd {
		case "forward":
			pos += val
		case "up":
			depth -= val
		case "down":
			depth += val
		}
		//fmt.Printf("cmd: %v, %v = %v\n", i, cmd, val)
	}

	return pos * depth
}
