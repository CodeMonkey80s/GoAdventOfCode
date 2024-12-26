package day8_part2

import (
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	ids := make(map[int]string)
	for id, line := range lines {
		switch {
		case strings.HasPrefix(line, "jmp"):
			ids[id] = "jmp"
		case strings.HasPrefix(line, "nop"):
			ids[id] = "nop"
		}
	}

	acc := -1
	valid := false
	for id, op := range ids {
		newLines := slices.Clone(lines)
		switch {
		case op == "jmp":
			newLines[id] = "nop" + lines[id][3:]
		case op == "nop":
			newLines[id] = "jmp" + lines[id][3:]
		}
		acc, valid = runProgram(newLines)
		if valid {
			break
		}
	}

	return acc
}

func runProgram(lines []string) (int, bool) {
	acc := 0
	valid := false

	cache := make(map[int]int)

	n := 0
loop:
	for {
		if n == len(lines) {
			valid = true
			break loop
		}

		line := lines[n]
		cache[n]++
		c, ok := cache[n]
		if ok && c >= 2 {
			valid = false
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
		if n >= len(lines) {
			valid = true
			break loop
		}
	}

	return acc, valid
}
