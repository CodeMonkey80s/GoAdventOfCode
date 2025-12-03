package day2_part1

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
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
		},
		Output: 1227775554,
	},
}

func Test_getAnswerForPart1(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day2_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 0,
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
