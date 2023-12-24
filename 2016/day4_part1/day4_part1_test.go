package day4_part1

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
			"aaaaa-bbb-z-y-x-123[abxyz]",
			"a-b-c-d-e-f-g-h-987[abcde]",
			"not-a-real-room-404[oarel]",
			"totally-real-room-200[decoy]",
		},
		Output: 1514,
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
			Output: 409147,
		},
	}
	testCases = append(testCases, testCase...)
}
func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isValidPassword(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "aaaaa-bbb-z-y-x-123[abxyz]",
			Output: true,
		},
		{
			Input:  "a-b-c-d-e-f-g-h-987[abcde]",
			Output: true,
		},
		{
			Input:  "not-a-real-room-404[oarel]",
			Output: true,
		},
		{
			Input:  "totally-real-room-200[decoy]",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			_, output := isValidPassword(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
