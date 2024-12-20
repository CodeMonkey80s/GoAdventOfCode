package day12_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"R90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R90",
				"R90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R90",
				"R90",
				"R90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R90",
				"R90",
				"R90",
				"R90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R90",
				"R90",
				"R90",
				"R90",
				"R90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L90",
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L90",
				"L90",
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L90",
				"L90",
				"L90",
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L90",
				"L90",
				"L90",
				"L90",
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R90",
				"L90",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"R180",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"L180",
				"F10",
			},
			Output: 110,
		},
		{
			Input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			Output: 286,
		},
		{
			Input: []string{
				"F10",
				"N3",
				"F7",
				"L90",
				"F11",
			},
			Output: 274,
		},
	}
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 107281,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
