package day12_part2

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
				"cpy 41 a",
				"inc a",
				"inc a",
				"dec a",
				"jnz a 2",
				"dec a",
			},
			Output: 42,
		},
	}
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 9227647,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
