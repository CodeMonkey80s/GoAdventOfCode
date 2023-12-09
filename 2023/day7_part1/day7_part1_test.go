package day7_part1

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
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		},
		Output: 6440,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 251216224,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_totalWinnings(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := totalWinnings(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
