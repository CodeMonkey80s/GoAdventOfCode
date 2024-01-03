package day5_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"FBFBBFFRLR",
			},
			Output: 357,
		},
		{
			Input: []string{
				"BFFFBBFRRR",
			},
			Output: 567,
		},
		{
			Input: []string{
				"FFFBBBFRRR",
			},
			Output: 119,
		},
		{
			Input: []string{
				"BBFFBBFRLL",
			},
			Output: 820,
		},
	}
	lines := util.LoadInputFile("../inputs/day5_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 885,
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
