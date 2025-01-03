package day9_part1

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
			Output: 120765,
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
		Output string
		Length int
	}{
		{
			Input:  "ADVENT",
			Output: "ADVENT",
			Length: 6,
		},
		{
			Input:  "A(1x5)BC",
			Output: "ABBBBBC",
			Length: 7,
		},
		{
			Input:  "(3x3)XYZ",
			Output: "XYZXYZXYZ",
			Length: 9,
		},
		{
			Input:  "A(2x2)BCD(2x2)EFG",
			Output: "ABCBCDEFEFG",
			Length: 11,
		},
		{
			Input:  "(6x1)(1x3)A",
			Output: "(1x3)A",
			Length: 6,
		},
		{
			Input:  "X(8x2)(3x3)ABCY",
			Output: "X(3x3)ABC(3x3)ABCY",
			Length: 18,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v_%v\n", tc.Input, tc.Output, tc.Length)
		t.Run(label, func(t *testing.T) {
			output := decompressString(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
