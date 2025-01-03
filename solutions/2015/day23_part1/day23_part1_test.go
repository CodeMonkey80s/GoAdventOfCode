package day23_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input    []string
		Register rune
		Output   int
	}{
		{
			Input: []string{
				"inc a",
				"jio a, +2",
				"tpl a",
				"inc a",
			},
			Register: 'a',
			Output:   2,
		},
	}
	lines := util.LoadInputFile("../inputs/day23_input.txt")
	testCase := []struct {
		Input    []string
		Register rune
		Output   int
	}{
		{
			Input:    lines,
			Register: 'b',
			Output:   184,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			registers := getAnswer(tc.Input)
			output := registers[tc.Register]
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
