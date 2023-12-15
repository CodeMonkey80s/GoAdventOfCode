package day2_part1

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
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		},
		Output: 2,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day2_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: 586,
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
