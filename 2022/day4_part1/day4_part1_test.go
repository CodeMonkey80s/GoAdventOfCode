package day4_part1

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"2-4,6-8",
			"2-3,4-5",
			"5-7,7-9",
			"2-8,3-7",
			"6-6,4-6",
			"2-6,4-8",
		},
		Output: 2,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 556,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_calculateSum(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := pairsFullyContained(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_extractRanges(t *testing.T) {
	testCases := []struct {
		Input  string
		Output []Range
	}{
		{
			Input: "2-4,6-8",
			Output: []Range{
				{
					Lo: 2,
					Hi: 4,
				},
				{
					Lo: 6,
					Hi: 8,
				},
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := extractRanges(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doRangesOverlap(t *testing.T) {
	testCases := []struct {
		Input  []Range
		Output bool
	}{
		{
			Input: []Range{
				{
					Lo: 2,
					Hi: 8,
				},
				{
					Lo: 3,
					Hi: 7,
				},
			},
			Output: true,
		},
		{
			Input: []Range{
				{
					Lo: 6,
					Hi: 6,
				},
				{
					Lo: 4,
					Hi: 6,
				},
			},
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doRangesOverlap(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
