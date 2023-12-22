package day9_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"London to Dublin = 464",
			"London to Belfast = 518",
			"Dublin to Belfast = 141",
		},
		Output: 982,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 898,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getLongestRoute(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getLongestRoute(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
