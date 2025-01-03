package day6_part2

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
			"turn on 0,0 through 0,0",
		},
		Output: 1,
	},
	{
		Input: []string{
			"toggle 0,0 through 999,999",
		},
		Output: 2_000_000,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines,
			Output: 14687245,
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

func Test_getCommand(t *testing.T) {
	var testCases = []struct {
		Input     string
		OutputCMD string
		OutputPos Coords
	}{
		{
			Input:     "toggle 205,417 through 703,826",
			OutputCMD: "toggle",
			OutputPos: Coords{
				SX: 205,
				SY: 417,
				EX: 703,
				EY: 826,
			},
		},
		{
			Input:     "turn off 399,680 through 861,942",
			OutputCMD: "turn off",
			OutputPos: Coords{
				SX: 399,
				SY: 680,
				EX: 861,
				EY: 942,
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v %v\n", tc.Input, tc.OutputCMD, tc.OutputPos)
		t.Run(label, func(t *testing.T) {
			cmd, pos := getCommand(tc.Input)
			if cmd != tc.OutputCMD {
				t.Errorf("Expected output to be %q but we got %q", tc.OutputCMD, cmd)
			}
			if pos != tc.OutputPos {
				t.Errorf("Expected output to be %q but we got %q", tc.OutputPos, pos)
			}
		})
	}
}
