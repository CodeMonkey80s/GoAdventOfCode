package day11_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output int
	}{
		{
			Input:  "ne,ne,ne",
			Output: 3,
		},
		{
			Input:  "ne,ne,sw,sw",
			Output: 0,
		},
		{
			Input:  "ne,ne,s,s",
			Output: 2,
		},
		{
			Input:  "se,sw,se,sw,sw",
			Output: 3,
		},
	}
	lines := util.LoadInputFile("../inputs/day11_input.txt")
	testCase := []struct {
		Input  string
		Output int
	}{
		{
			Input:  lines[0],
			Output: 824,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
