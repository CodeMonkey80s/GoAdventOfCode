package day18_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Steps  int
		Output int
	}{
		{
			Input: []string{
				"##.#.#",
				"...##.",
				"#....#",
				"..#...",
				"#.#..#",
				"####.#",
			},
			Steps:  5,
			Output: 17,
		},
	}
	lines := util.LoadInputFile("../inputs/day18_input.txt")
	testCase := []struct {
		Input  []string
		Steps  int
		Output int
	}{
		{
			Input:  lines,
			Steps:  100,
			Output: 1006,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.Steps)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
