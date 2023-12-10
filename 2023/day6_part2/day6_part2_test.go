package day6_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/2023/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"Time:      7  15   30",
			"Distance:  9  40  200",
		},
		Output: 71503,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 27363861,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getTheNumber(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getTheNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
