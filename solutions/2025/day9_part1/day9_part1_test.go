package day9_part1

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
			"7,1",
			"11,1",
			"11,7",
			"9,7",
			"9,5",
			"2,5",
			"2,3",
			"7,3",
		},
		Output: 50,
	},
}

func Test_getAnswerForPart1(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 4750297200,
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
