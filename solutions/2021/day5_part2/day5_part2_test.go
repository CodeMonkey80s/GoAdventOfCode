package day5_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  []string
	Output int
}{
	{
		Name: "Example Input",
		Input: []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		},
		Output: 12,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day5_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: 19472,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_howManyPointsOverlap(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := howManyPointsOverlap(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
