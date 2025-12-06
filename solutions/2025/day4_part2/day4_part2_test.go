package day4_part2

import (
	"GoAdventOfCode/util"
	"fmt"
	"testing"
)

var testCasesForPart2 = []struct {
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
		Output: 43,
	},
}

func Test_getAnswerForPart2(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 8184,
		},
	}
	testCasesForPart2 = append(testCasesForPart2, testCase...)
	for _, tc := range testCasesForPart2 {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswerForPart2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
