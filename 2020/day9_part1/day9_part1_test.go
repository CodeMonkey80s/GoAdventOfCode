package day9_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: []string{
				"35",
				"20",
				"15",
				"25",
				"47",
				"40",
				"62",
				"55",
				"65",
				"95",
				"102",
				"117",
				"150",
				"182",
				"127",
				"219",
				"299",
				"277",
				"309",
				"576",
			},
			InputB: 5,
			Output: 127,
		},
	}
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCase := []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: lines,
			InputB: 25,
			Output: 1124361034,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
