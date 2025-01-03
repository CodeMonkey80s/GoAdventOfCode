package day10_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input       []string
		InputCycles int
		Output      int
		OutputX     int
	}{
		// {
		// 	Input: []string{
		// 		"noop   ",
		// 		"addx 3 ",
		// 		"addx -5",
		// 	},
		// 	InputCycles: 5,
		// 	Output:      0,
		// 	OutputX:     -1,
		// },
		// {
		// 	Input: []string{
		// 		"addx 15",
		// 		"addx -11",
		// 		"addx 6",
		// 		"addx -3",
		// 		"addx 5",
		// 		"addx -1",
		// 		"addx -8",
		// 		"addx 13",
		// 		"addx 4",
		// 		"noop",
		// 	},
		// 	InputCycles: 20,
		// 	Output:      0,
		// 	OutputX:     21,
		// },
		{
			Input: []string{
				"addx 15",
				"addx -11",
				"addx 6",
				"addx -3",
				"addx 5",
				"addx -1",
				"addx -8",
				"addx 13",
				"addx 4",
				"noop",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx 5",
				"addx -1",
				"addx -35",
				"addx 1",
				"addx 24",
				"addx -19",
				"addx 1",
				"addx 16",
				"addx -11",
				"noop",
				"noop",
				"addx 21",
				"addx -15",
				"noop",
				"noop",
				"addx -3",
				"addx 9",
				"addx 1",
				"addx -3",
				"addx 8",
				"addx 1",
				"addx 5",
				"noop",
				"noop",
				"noop",
				"noop",
				"noop",
				"addx -36",
				"noop",
				"addx 1",
				"addx 7",
				"noop",
				"noop",
				"noop",
				"addx 2",
				"addx 6",
				"noop",
				"noop",
				"noop",
				"noop",
				"noop",
				"addx 1",
				"noop",
				"noop",
				"addx 7",
				"addx 1",
				"noop",
				"addx -13",
				"addx 13",
				"addx 7",
				"noop",
				"addx 1",
				"addx -33",
				"noop",
				"noop",
				"noop",
				"addx 2",
				"noop",
				"noop",
				"noop",
				"addx 8",
				"noop",
				"addx -1",
				"addx 2",
				"addx 1",
				"noop",
				"addx 17",
				"addx -9",
				"addx 1",
				"addx 1",
				"addx -3",
				"addx 11",
				"noop",
				"noop",
				"addx 1",
				"noop",
				"addx 1",
				"noop",
				"noop",
				"addx -13",
				"addx -19",
				"addx 1",
				"addx 3",
				"addx 26",
				"addx -30",
				"addx 12",
				"addx -1",
				"addx 3",
				"addx 1",
				"noop",
				"noop",
				"noop",
				"addx -9",
				"addx 18",
				"addx 1",
				"addx 2",
				"noop",
				"noop",
				"addx 9",
				"noop",
				"noop",
				"noop",
				"addx -1",
				"addx 2",
				"addx -37",
				"addx 1",
				"addx 3",
				"noop",
				"addx 15",
				"addx -21",
				"addx 22",
				"addx -6",
				"addx 1",
				"noop",
				"addx 2",
				"addx 1",
				"noop",
				"addx -10",
				"noop",
				"noop",
				"addx 20",
				"addx 1",
				"addx 2",
				"addx 2",
				"addx -6",
				"addx -11",
				"noop",
				"noop",
				"noop",
			},
			Output:      0,
			OutputX:     17,
			InputCycles: 240,
		},
	}
	lines := util.LoadInputFile("../inputs/day10_input.txt")
	testCase := []struct {
		Input       []string
		InputCycles int
		Output      int
		OutputX     int
	}{
		{
			Input:       lines,
			Output:      0,
			OutputX:     22,
			InputCycles: 221,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output, x := getAnswer(tc.Input, tc.InputCycles)
			if x != tc.OutputX {
				t.Errorf("Expected output of x to be '%d' but we got '%d'", tc.OutputX, x)
			}
			if tc.Output > 0 {
				if output != tc.Output {
					t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
				}
			}
		})
	}
}
