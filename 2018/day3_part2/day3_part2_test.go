package day3_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  []string
	Output string
}{
	{
		Input: []string{
			"#1 @ 1,3: 4x4",
			"#2 @ 3,1: 4x4",
			"#3 @ 5,5: 2x2",
		},
		Output: "#3",
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day3_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output string
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: "#346",
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
