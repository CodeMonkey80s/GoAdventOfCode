package day2_part1

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
			"A Y",
			"B X",
			"C Z",
		},
		Output: 15,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day2_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 10816,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_calculateScore(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := calculateScore(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_roundScore(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []byte
		Output int
	}{
		// Draw
		{
			Name:   "Rock vs Rock",
			Input:  []byte{'A', 'X'},
			Output: 1 + 3,
		},
		{
			Name:   "Paper vs Paper",
			Input:  []byte{'B', 'Y'},
			Output: 2 + 3,
		},
		{
			Name:   "Scissors vs Scissors",
			Input:  []byte{'C', 'Z'},
			Output: 3 + 3,
		},
		// Win
		{
			Name:   "Scissors vs Rock",
			Input:  []byte{'C', 'X'},
			Output: 1 + 6,
		},
		{
			Name:   "Rock vs Paper",
			Input:  []byte{'A', 'Y'},
			Output: 2 + 6,
		},
		{
			Name:   "Paper vs Scissors",
			Input:  []byte{'B', 'Z'},
			Output: 3 + 6,
		},
		// Loss
		{
			Name:   "Rock vs Scissors",
			Input:  []byte{'A', 'Z'},
			Output: 3 + 0,
		},
		{
			Name:   "Paper vs Rock",
			Input:  []byte{'B', 'X'},
			Output: 1 + 0,
		},
		{
			Name:   "Scissors vs Paper",
			Input:  []byte{'C', 'Y'},
			Output: 2 + 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := roundScore(tt.Input[0], tt.Input[1]); got != tt.Output {
				t.Errorf("roundScore() = %v, but should be %v", got, tt.Output)
			}
		})
	}
}
