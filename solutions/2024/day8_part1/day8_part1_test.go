package day8_part1

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
				"..........",
				"..........",
				"..........",
				"....a.....",
				"..........",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			Output: 2,
		},
		{
			Input: []string{
				"..........",
				"..........",
				"..........",
				"....a.....",
				"........a.",
				".....a....",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			Output: 4,
		},
		{
			Input: []string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			Output: 14,
		},
	}
	lines := util.LoadInputFile("../inputs/day8_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 396,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
