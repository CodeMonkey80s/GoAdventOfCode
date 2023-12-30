package day8_part2

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		InputW int
		InputH int
		Output []string
	}{
		{
			Input:  "0222112222120000",
			InputW: 2,
			InputH: 2,
			Output: []string{
				"01",
				"01",
			},
		},
	}
	lines := util.LoadInputFile("../inputs/day8_input.txt")
	testCase := []struct {
		Input  string
		InputW int
		InputH int
		Output []string
	}{
		{
			Input:  lines[0],
			InputW: 25,
			InputH: 6,
			Output: []string{
				"xx",
				"yy",
			},
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.InputW, tc.InputH)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
