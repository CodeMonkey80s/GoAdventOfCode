package day18_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input      string
		InputRows  int
		Output     []string
		OutputSafe int
	}{
		{
			Input:     "..^^.",
			InputRows: 3,
			Output: []string{
				"..^^.",
				".^^^^",
				"^^..^",
			},
			OutputSafe: 6,
		},
		{
			Input:     ".^^.^.^^^^",
			InputRows: 10,
			Output: []string{
				".^^.^.^^^^",
				"^^^...^..^",
				"^.^^.^.^^.",
				"..^^...^^^",
				".^^^^.^^.^",
				"^^..^.^^..",
				"^^^^..^^^.",
				"^..^^^^.^^",
				".^^^..^.^^",
				"^^.^^^..^^",
			},
			OutputSafe: 38,
		},
	}
	lines := util.LoadInputFile("../inputs/day18_input.txt")
	testCase := []struct {
		Input      string
		InputRows  int
		Output     []string
		OutputSafe int
	}{
		{
			Input:      lines[0],
			InputRows:  40,
			Output:     []string{""},
			OutputSafe: 1951,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			_, safe := getOutput(tc.Input, tc.InputRows)
			if safe != tc.OutputSafe {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputSafe, safe)
			}
		})
	}
}
