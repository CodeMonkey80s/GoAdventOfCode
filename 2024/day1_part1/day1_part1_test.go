package day1_part1

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
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		},
		Output: 11,
	},
}

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 2086478,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
