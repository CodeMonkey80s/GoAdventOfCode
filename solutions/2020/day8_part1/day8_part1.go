package day8_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	acc := 0

	cache := make(map[int]int)

	n := 0
loop:
	for {
		line := lines[n]
		cache[n]++
		c, ok := cache[n]
		if ok && c >= 2 {
			break loop
		}

		parts := strings.Fields(line)
		cmd := parts[0]
		// fmt.Printf("n: %d, cmd: %q, parts: %q\n", n, cmd, parts)

		switch cmd {
		case "nop":
		case "acc":
			acc += util.ConvertStringToInt(parts[1])
		case "jmp":
			offset := util.ConvertStringToInt(parts[1])
			n += offset
			continue loop
		}

		n++
	}

	return acc
}
