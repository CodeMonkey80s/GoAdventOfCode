package day6_part1

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
			"123 328  51 64",
			"45 64  387 23",
			"6 98  215 314",
			"*   +   *   +",
		},
		Output: 4277556,
	},
}

func Test_getAnswerForPart1(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 5595593539811,
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
