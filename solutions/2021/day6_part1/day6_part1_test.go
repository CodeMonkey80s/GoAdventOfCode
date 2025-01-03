package day6_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  string
	Output int
}{
	{
		Name:   "Example Input",
		Input:  "3,4,3,1,2",
		Output: 5934,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Name   string
		Input  string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines[0],
			Output: 351188,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_howManyPointsOverlap(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := howManyLanternfish(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
