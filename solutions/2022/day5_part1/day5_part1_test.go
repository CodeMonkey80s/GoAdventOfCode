package day5_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output string
}{
	{
		Input: []string{
			"    [D]           ",
			"[N] [C]           ",
			"[Z] [M] [P]       ",
			" 1   2   3        ",
			"",
			"move 1 from 2 to 1",
			"move 3 from 1 to 3",
			"move 2 from 2 to 1",
			"move 1 from 1 to 2",
		},
		Output: "CMZ",
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day5_input.txt")
	testCase := []struct {
		Input  []string
		Output string
	}{
		{
			Input:  lines,
			Output: "TBVFVDZPN",
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_calculateSum(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := finalCratesPositions(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
