package day1_part1

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
			"199",
			"200",
			"208",
			"210",
			"200",
			"207",
			"240",
			"269",
			"260",
			"263",
		},
		Output: 7,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: 1602,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_startOfPacket(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := howManyMeasurementsLarger(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
