package day1_part2

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
		Input: []string{
			"+1",
			"-2",
			"+3",
			"+1",
			"+1",
			"-2",
			"+3",
			"+1",
		},
		Output: 2,
	},
	{
		Input: []string{
			"+1",
			"-1",
		},
		Output: 0,
	},
	{
		Input: []string{
			"+3",
			"+3",
			"+4",
			"-2",
			"-4",
		},
		Output: 10,
	},
	{
		Input: []string{
			"-6",
			"+3",
			"+8",
			"+5",
			"-6",
		},
		Output: 5,
	},
	{
		Input: []string{
			"+7",
			"+7",
			"-2",
			"-7",
			"-4",
		},
		Output: 14,
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
			Output: 488,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := ""
		if tc.Name != "" {
			label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		}
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
