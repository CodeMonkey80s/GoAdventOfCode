package day1_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  string
	Output int
}{
	{
		Input:  "R2, L3",
		Output: 5,
	},
	{
		Input:  "R2, R2, R2",
		Output: 2,
	},
	{
		Input:  "R5, L5, R5, R3",
		Output: 12,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Name   string
		Input  string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines[0],
			Output: 271,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_Player_updateFacing(t *testing.T) {
	var testCases = []struct {
		Initial rune
		Input   rune
		Output  rune
	}{
		{
			Initial: 'N',
			Input:   'L',
			Output:  'W',
		},
		{
			Initial: 'N',
			Input:   'R',
			Output:  'E',
		},
		{
			Initial: 'S',
			Input:   'R',
			Output:  'W',
		},
		{
			Initial: 'S',
			Input:   'L',
			Output:  'E',
		},
		{
			Initial: 'E',
			Input:   'L',
			Output:  'N',
		},
		{
			Initial: 'E',
			Input:   'R',
			Output:  'S',
		},
		{
			Initial: 'W',
			Input:   'R',
			Output:  'N',
		},
		{
			Initial: 'W',
			Input:   'L',
			Output:  'S',
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %c %c Output: %c\n", tc.Initial, tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			p := Player{
				Facing: tc.Initial,
			}
			p.updateFacing(tc.Input)
			if p.Facing != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, p.Facing)
			}
		})
	}
}
