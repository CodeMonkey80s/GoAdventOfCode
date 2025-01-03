package day6_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCases := []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"abc",
			},
			Output: 3,
		},
		{
			Input: []string{
				"a",
				"b",
				"c",
			},
			Output: 3,
		},
		{
			Input: []string{
				"ab",
				"ac",
			},
			Output: 3,
		},
		{
			Input: []string{
				"a",
				"a",
				"a",
				"a",
			},
			Output: 1,
		},
		{
			Input: []string{
				"abc",
				"",
				"a",
				"b",
				"c",
				"",
				"ab",
				"ac",
				"",
				"a",
				"a",
				"a",
				"a",
				"",
				"b",
			},
			Output: 11,
		},
		{
			Input:  lines,
			Output: 6310,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
