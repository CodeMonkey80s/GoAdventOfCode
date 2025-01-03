package day6_part2

import (
	"fmt"
	"strings"
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
			Output: 0,
		},
		{
			Input: []string{
				"ab",
				"ac",
			},
			Output: 1,
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
				"b",
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
			Output: 6,
		},
		{
			Input:  lines,
			Output: 3193,
		},
	}
	for _, tc := range testCases {
		name := strings.Join(tc.Input, "")
		if len(tc.Input) > 20 {
			name = "Puzzle Input"
		}
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
