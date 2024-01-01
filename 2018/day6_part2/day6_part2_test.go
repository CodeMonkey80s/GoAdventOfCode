package day6_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input    []string
		InputMax int
		Output   int
	}{
		{
			Input: []string{
				"1, 1",
				"1, 6",
				"8, 3",
				"3, 4",
				"5, 5",
				"8, 9",
			},
			InputMax: 32,
			Output:   16,
		},
	}
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Input    []string
		InputMax int
		Output   int
	}{
		{
			Input:    lines,
			InputMax: 10_000,
			Output:   42998,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.InputMax)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
