package day8_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		InputW int
		InputH int
		Output int
	}{
		{
			Input:  "123456789012",
			InputW: 3,
			InputH: 2,
			Output: 1,
		},
	}
	lines := util.LoadInputFile("../inputs/day8_input.txt")
	testCase := []struct {
		Input  string
		InputW int
		InputH int
		Output int
	}{
		{
			Input:  lines[0],
			InputW: 25,
			InputH: 6,
			Output: 1360,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.InputW, tc.InputH)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
