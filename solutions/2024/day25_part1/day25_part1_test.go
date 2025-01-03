package day25_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{

		{
			Input: []string{
				"#####",
				".####",
				".####",
				".####",
				".#.#.",
				".#...",
				".....",
				"",
				"#####",
				"##.##",
				".#.##",
				"...##",
				"...#.",
				"...#.",
				".....",
				"",
				".....",
				"#....",
				"#....",
				"#...#",
				"#.#.#",
				"#.###",
				"#####",
				"",
				".....",
				".....",
				"#.#..",
				"###..",
				"###.#",
				"###.#",
				"#####",
				"",
				".....",
				".....",
				".....",
				"#....",
				"#.#..",
				"#.#.#",
				"#####",
				"",
			},
			Output: 3,
		},
	}
	lines := util.LoadInputFile("../inputs/day25_input.txt")
	lines = append(lines, "\n")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 3065,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
