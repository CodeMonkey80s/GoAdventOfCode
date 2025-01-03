package day8_part1

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
			`""`,
			`"abc"`,
			`"aaa\"aaa"`,
			`"\x27"`,
		},
		Output: 12,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day8_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 1342,
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

func Test_countBytesAndRunes(t *testing.T) {
	var testCases = []struct {
		Input       string
		OutputBytes int
		OutputRunes int
	}{
		{
			Input:       `""`,
			OutputBytes: 2,
			OutputRunes: 0,
		},
		{
			Input:       `"abc"`,
			OutputBytes: 5,
			OutputRunes: 3,
		},
		{
			Input:       `"aaa\"aaa"`,
			OutputBytes: 10,
			OutputRunes: 7,
		},
		{
			Input:       `"\x27"`,
			OutputBytes: 6,
			OutputRunes: 1,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v %v\n", len(tc.Input), tc.OutputBytes, tc.OutputRunes)
		t.Run(label, func(t *testing.T) {
			bytes, runes := countBytesAndRunes(tc.Input)
			if bytes != tc.OutputBytes || runes != tc.OutputRunes {
				t.Errorf("Expected output to be %v %v but we got %v %v", tc.OutputBytes, tc.OutputRunes, bytes, runes)
			}
		})
	}
}
