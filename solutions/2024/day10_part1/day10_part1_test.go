package day10_part1

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
				"0123",
				"1234",
				"8765",
				"9876",
			},
			Output: 1,
		},
		{
			Input: []string{
				"...0...",
				"...1...",
				"...2...",
				"6543456",
				"7.....7",
				"8.....8",
				"9.....9",
			},
			Output: 2,
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
			Output: 4,
		},
		{
			Input: []string{
				"10..9..",
				"2...8..",
				"3...7..",
				"4567654",
				"...8..3",
				"...9..2",
				".....01",
			},
			Output: 3,
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
			Output: 36,
		},
	}
	lines := util.LoadInputFile("../inputs/day10_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 538,
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
