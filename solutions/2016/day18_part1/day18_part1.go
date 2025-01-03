package day18_part1

import (
	"strings"
)

func getOutput(input string, rows int) ([]string, int) {
	var output []string

	output = append(output, input)

	input = "s" + input + "s"

	var sb strings.Builder
	for i := 0; i < rows-1; i++ {
		a := 0
		b := 1
		c := 2
		for j := 0; j < len(input)-2; j++ {
			switch {
			case input[a] == '^' && input[b] == '^' && input[c] != '^':
				sb.WriteRune('^')
			case input[a] != '^' && input[b] == '^' && input[c] == '^':
				sb.WriteRune('^')
			case input[a] == '^' && input[b] != '^' && input[c] != '^':
				sb.WriteRune('^')
			case input[a] != '^' && input[b] != '^' && input[c] == '^':
				sb.WriteRune('^')
			default:
				sb.WriteRune('.')
			}
			a++
			b++
			c++
		}

		input = "s" + sb.String() + "s"
		output = append(output, sb.String())
		sb.Reset()
	}

	safe := 0
	for _, line := range output {
		for _, ch := range line {
			if ch == '.' {
				safe++
			}
		}
	}

	// util.PrintSlice("output", output)

	return output, safe
}
