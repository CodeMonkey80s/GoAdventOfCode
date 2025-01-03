package day9_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCases := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 11658395076,
		},
	}
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

func Test_decompressString(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output int
	}{
		{
			Input:  "(3x3)XYZ",
			Output: 9,
		},
		{
			Input:  "X(8x2)(3x3)ABCY",
			Output: 20,
		},
		{
			Input:  "(27x12)(20x12)(13x14)(7x10)(1x12)A",
			Output: 241920,
		},
		{
			Input:  "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			Output: 445,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := decompressString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
