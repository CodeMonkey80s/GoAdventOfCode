package day14_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA []string
		InputB int
		InputC int
		InputD int
		Output int
	}{
		{
			InputA: []string{
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-3",
				"p=10,3 v=-1,2",
				"p=2,0 v=2,-1",
				"p=0,0 v=1,3",
				"p=3,0 v=-2,-2",
				"p=7,6 v=-1,-3",
				"p=3,0 v=-1,-2",
				"p=9,3 v=2,3",
				"p=7,3 v=-1,2",
				"p=2,4 v=2,-3",
				"p=9,5 v=-3,-3",
			},
			InputB: 11,
			InputC: 7,
			InputD: 100,
			Output: 12,
		},
	}
	lines := util.LoadInputFile("../inputs/day14_input.txt")
	testCase := []struct {
		InputA []string
		InputB int
		InputC int
		InputD int
		Output int
	}{
		{
			InputA: lines,
			InputB: 101,
			InputC: 103,
			InputD: 100,
			Output: 225943500,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputA, tc.InputB, tc.InputC, tc.InputD)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
