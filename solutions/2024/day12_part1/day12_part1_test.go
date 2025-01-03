package day12_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getTotalPrice(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			Output: 140,
		},
		{
			Input: []string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			Output: 1930,
		},
	}
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 1488414,
		},
	}
	testCases = append(testCases, testCase...)
	for i, tc := range testCases {
		label := fmt.Sprintf("%v_%d\n", "Puzzle Input", i)
		t.Run(label, func(t *testing.T) {
			output := getTotalPrice(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be of length '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
