package day5_part1

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
			"ugknbfddgicrmopn",
		},
		Output: 1,
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
			Output: -1,
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
			Input:  "ugknbfddgicrmopn",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: true,
		},
		{
			Input:  "jchzalrnumimnmhp",
			Output: false,
		},
		{
			Input:  "haegwjzuvuyypxyu",
			Output: false,
		},
		{
			Input:  "dvszwmarrgswjxmb",
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

func Test_doesStringContainThreeVowels(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "aeiouaeiouaeiou",
			Output: true,
		},
		{
			Input:  "ugknbfddgicrmopn",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: true,
		},
		{
			Input:  "jchzalrnumimnmhp",
			Output: true,
		},
		{
			Input:  "haegwjzuvuyypxyu",
			Output: true,
		},
		{
			Input:  "dvszwmarrgswjxmb",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doesStringContainThreeVowels(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doesStringContainDoubleLetter(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "ugknbfddgicrmopn",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: true,
		},
		{
			Input:  "jchzalrnumimnmhp",
			Output: false,
		},
		{
			Input:  "haegwjzuvuyypxyu",
			Output: true,
		},
		{
			Input:  "dvszwmarrgswjxmb",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doesStringContainDoubleLetter(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doesStringContainBadWord(t *testing.T) {
	var testCases = []struct {
		Name   string
		Input  string
		Output bool
	}{
		{
			Input:  "jchzalrnumimnmhp",
			Output: false,
		},
		{
			Input:  "haegwjzuvuyypxyu",
			Output: true,
		},
		{
			Input:  "dvszwmarrgswjxmb",
			Output: false,
		},
		{
			Input:  "dvszwmabrgswjxmb",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doesStringContainBadWord(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
