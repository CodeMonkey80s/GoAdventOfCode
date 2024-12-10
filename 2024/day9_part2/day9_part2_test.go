package day9_part2

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_unpack(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output []int
	}{
		{
			Input:  "12345",
			Output: []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2},
		},
		{
			Input:  "909090",
			Output: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Input)
		t.Run(label, func(t *testing.T) {
			output := unpackDiskMap(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output int
	}{
		{
			Input:  "2333133121414131402",
			Output: 2858,
		},
	}
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCase := []struct {
		Input  string
		Output int
	}{
		{
			Input:  lines[0],
			Output: 6421724645083,
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
