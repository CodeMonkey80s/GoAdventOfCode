package day23_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) map[rune]int {

	registers := map[rune]int{
		'a': 1,
		'b': 0,
	}

	n := 0
loop:
	for {
		line := lines[n]
		parts := strings.Fields(line)
		cmd := parts[0]
		arg := rune(parts[1][0])

		switch cmd {
		case "inc":
			registers[arg] += 1
		case "tpl":
			registers[arg] *= 3
		case "hlf":
			registers[arg] /= 2
		case "jmp":
			offset := util.ConvertStringToInt(parts[1])
			n += offset
			goto check
		case "jie":
			if registers[arg]%2 == 0 {
				offset := util.ConvertStringToInt(parts[2])
				n += offset
				goto check
			}
		case "jio":
			if registers[arg] == 1 {
				offset := util.ConvertStringToInt(parts[2])
				n += offset
				goto check
			}
		}
		n++
	check:
		if n >= len(lines) {
			break loop
		}
	}

	return registers
}
