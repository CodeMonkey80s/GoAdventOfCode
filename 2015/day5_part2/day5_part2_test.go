package day5_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  []string
	Output int
}{
	{
		Input: []string{
			"qjhvhtzxzqqjkmpb",
			"xxyxx",
		},
		Output: 2,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: 55,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input[0], tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isStringNice(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "qjhvhtzxzqqjkmpb",
			Output: true,
		},
		{
			Input:  "xxyxx",
			Output: true,
		},
		{
			Input:  "uurcxstgmygtbstg",
			Output: false,
		},
		{
			Input:  "ieodomkazucvgmuy",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isStringNice(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doesStringContainTwoLettersAppearTwice(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "qjhvhtzxzqqjkmpb",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: false,
		},
		{
			Input:  "xxyxx",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doesStringContainPairOfAnyTwoLetters(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doesStringContainLetterThatRepeats(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "xyx",
			Output: true,
		},
		{
			Input:  "urctgmystg",
			Output: false,
		},
		{
			Input:  "abcdefeghi",
			Output: true,
		},
		{
			Input:  "efe",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doesStringContainLetterThatRepeats(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
