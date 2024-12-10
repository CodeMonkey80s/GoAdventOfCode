package day10_part2

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
				".....0.",
				"..4321.",
				"..5..2.",
				"..6543.",
				"..7..4.",
				"..8765.",
				"..9....",
			},
			Output: 3,
		},
		{
			Input: []string{
				"..90..9",
				"...1.98",
				"...2..7",
				"6543456",
				"765.987",
				"876....",
				"987....",
			},
			Output: 13,
		},
		{
			Input: []string{
				"012345",
				"123456",
				"234567",
				"345678",
				"4.6789",
				"56789.",
			},
			Output: 227,
		},
		{
			Input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			Output: 81,
		},
	}
	lines := util.LoadInputFile("../inputs/day10_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 1110,
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
