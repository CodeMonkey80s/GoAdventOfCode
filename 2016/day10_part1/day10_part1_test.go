package day10_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input         []string
		CompareValue1 int
		CompareValue2 int
		Output        int
	}{
		{
			Input: []string{
				"value 5 goes to bot 2",
				"bot 2 gives low to bot 1 and high to bot 0",
				"value 3 goes to bot 1",
				"bot 1 gives low to output 1 and high to bot 0",
				"bot 0 gives low to output 2 and high to output 0",
				"value 2 goes to bot 2",
			},
			CompareValue1: 5,
			CompareValue2: 2,
			Output:        2,
		},
	}
	lines := util.LoadInputFile("../inputs/day10_input.txt")
	testCase := []struct {
		Input         []string
		CompareValue1 int
		CompareValue2 int
		Output        int
	}{
		{
			Input:         lines,
			CompareValue1: 61,
			CompareValue2: 17,
			Output:        98,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.CompareValue1, tc.CompareValue2)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
