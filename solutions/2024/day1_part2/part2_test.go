package day1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCasesForPart2 = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		},
		Output: 31,
	},
}

func Test_getAnswerForPart2(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 24941624,
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
