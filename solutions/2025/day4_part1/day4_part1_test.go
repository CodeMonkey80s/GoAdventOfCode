package day4_part1

import (
	"GoAdventOfCode/util"
	"fmt"
	"testing"
)

var testCasesForPart1 = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"..@@.@@@@.",
			"@@@.@.@.@@",
			"@@@@@.@.@@",
			"@.@@@@..@.",
			"@@.@@@@.@@",
			".@@@@@@@.@",
			".@.@.@.@@@",
			"@.@@@.@@@@",
			".@@@@@@@@.",
			"@.@.@@@.@.",
		},
		Output: 13,
	},
}

func Test_getAnswerForPart1(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 1,
		},
	}
	testCasesForPart1 = append(testCasesForPart1, testCase...)
	for _, tc := range testCasesForPart1 {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswerForPart1(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
