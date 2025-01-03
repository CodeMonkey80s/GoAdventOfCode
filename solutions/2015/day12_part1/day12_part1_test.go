package day12_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  `[1,2,3]`,
		Output: 6,
	},
	{
		Input:  `{"a":2,"b":4}`,
		Output: 6,
	},
	{
		Input:  `[[[3]]]`,
		Output: 3,
	},
	{
		Input:  `{"a":{"b":4},"c":-1}`,
		Output: 3,
	},
	{
		Input:  `[]`,
		Output: 0,
	},
	{
		Input:  `{}`,
		Output: 0,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  string
		Output int
	}{
		{
			Input:  lines[0],
			Output: 191164,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
