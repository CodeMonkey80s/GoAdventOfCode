package day12_part2

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
			Output: 80,
		},
		{
			Input: []string{
				"EEEEE",
				"EXXXX",
				"EEEEE",
				"EXXXX",
				"EEEEE",
			},
			Output: 236,
		},
		{
			Input: []string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			Output: 436,
		},
		{
			Input: []string{
				"AAAAAA",
				"AAABBA",
				"AAABBA",
				"ABBAAA",
				"ABBAAA",
				"AAAAAA",
			},
			Output: 368,
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
			Output: 1206,
		},
	}
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 911750,
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
