package day4_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCases := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 991,
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

func Test_decryptPassword(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output string
	}{
		{
			Input:  "qzmt-zixmtkozy-ivhz-343[abcdf]",
			Output: "very encrypted name",
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%q_%q\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			_, output := decryptPassword(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}
