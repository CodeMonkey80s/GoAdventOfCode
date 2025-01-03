package day1_part2

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
		Input:  "1212",
		Output: 6,
	},
	{
		Input:  "1221",
		Output: 0,
	},
	{
		Input:  "123425",
		Output: 4,
	},
	{
		Input:  "123123",
		Output: 12,
	},
	{
		Input:  "12131415",
		Output: 4,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Name   string
		Input  string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines[0],
			Output: 1244,
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
