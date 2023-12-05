package day3_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/2023/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
		Output: 4361,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day3_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 554003,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_sumOfAllOfThePartNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfAllOfThePartNumbers(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
