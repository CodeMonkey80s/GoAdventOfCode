package day3_part1

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
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		},
		Output: 157,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day3_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 7811,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_calculateSum(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := calculateSum(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_valueFromCharacter(t *testing.T) {
	tests := []struct {
		Input  byte
		Output int
	}{
		{
			Input:  'a',
			Output: 1,
		},
		{
			Input:  'z',
			Output: 26,
		},
		{
			Input:  'A',
			Output: 27,
		},
		{
			Input:  'Z',
			Output: 52,
		},
	}
	for _, tc := range tests {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			if got := valueFromCharacter(tc.Input); got != tc.Output {
				t.Errorf("roundScore() = %v, but should be %v", got, tc.Output)
			}
		})
	}
}
