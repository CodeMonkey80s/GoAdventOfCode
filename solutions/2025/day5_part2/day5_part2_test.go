package day5_part2

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
			"3-5",
			"10-14",
			"16-20",
			"12-18",
		},
		Output: 14,
	},
}

func Test_getAnswerForPart2(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day5_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 1,
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
