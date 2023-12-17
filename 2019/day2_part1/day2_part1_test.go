package day2_part1

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
		Input:  "1,9,10,3,2,3,11,0,99,30,40,50",
		Output: 3500,
	},
	{
		Input:  "1,0,0,0,99",
		Output: 2,
	},
	{
		Input:  "2,3,0,3,99",
		Output: 2,
	},
	{
		Input:  "2,4,4,5,99,0",
		Output: 2,
	},
	{
		Input:  "1,1,1,4,99,5,6,0,99",
		Output: 30,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day2_input.txt")
	testCase := []struct {
		Name   string
		Input  string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines[0],
			Output: 5866663,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
