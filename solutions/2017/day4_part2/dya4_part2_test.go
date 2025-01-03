package day4_part2

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
			"abcde fghij",
			"abcde xyz ecdab",
			"a ab abc abd abf abj",
			"iiii oiii ooii oooi oooo",
			"oiii ioii iioi iiio",
		},
		Output: 3,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 251,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isValidPassphrase(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "abcde fghij",
			Output: true,
		},
		{
			Input:  "abcde xyz ecdab",
			Output: false,
		},
		{
			Input:  "a ab abc abd abf abj",
			Output: true,
		},
		{
			Input:  "iiii oiii ooii oooi oooo",
			Output: true,
		},
		{
			Input:  "oiii ioii iioi iiio",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isValidPassphrase(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
