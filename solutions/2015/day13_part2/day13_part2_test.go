package day13_part2

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
			"Alice would gain 54 happiness units by sitting next to Bob.",
			"Alice would lose 79 happiness units by sitting next to Carol.",
			"Alice would lose 2 happiness units by sitting next to David.",
			"Bob would gain 83 happiness units by sitting next to Alice.",
			"Bob would lose 7 happiness units by sitting next to Carol.",
			"Bob would lose 63 happiness units by sitting next to David.",
			"Carol would lose 62 happiness units by sitting next to Alice.",
			"Carol would gain 60 happiness units by sitting next to Bob.",
			"Carol would gain 55 happiness units by sitting next to David.",
			"David would gain 46 happiness units by sitting next to Alice.",
			"David would lose 7 happiness units by sitting next to Bob.",
			"David would gain 41 happiness units by sitting next to Carol.",
		},
		Output: 286,
	},
}

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day13_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 601,
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
