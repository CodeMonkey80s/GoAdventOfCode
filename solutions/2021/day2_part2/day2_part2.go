package day2_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func horizontalPosition(commands []string) int {
	pos := 0
	depth := 0
	aim := 0
	for _, command := range commands {
		parts := strings.Fields(command)
		cmd := parts[0]
		val := util.ConvertStringToInt(parts[1])
		switch cmd {
		case "forward":
			pos += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}

	return pos * depth
}
