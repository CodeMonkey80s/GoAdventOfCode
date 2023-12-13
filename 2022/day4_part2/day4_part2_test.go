package day4_part2

import (
	"fmt"
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
		Output: 4,
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
			Output: 876,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_calculateSum(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := pairsOverlapping(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_doRangesOverlapAtAll(t *testing.T) {
	testCases := []struct {
		Input  []Range
		Output bool
	}{
		{
			Input: []Range{
				{
					Lo: 2,
					Hi: 4,
				},
				{
					Lo: 6,
					Hi: 8,
				},
			},
			Output: false,
		},
		{
			Input: []Range{
				{
					Lo: 5,
					Hi: 7,
				},
				{
					Lo: 7,
					Hi: 9,
				},
			},
			Output: true,
		},
		{
			Input: []Range{
				{
					Lo: 7,
					Hi: 9,
				},
				{
					Lo: 5,
					Hi: 7,
				},
			},
			Output: true,
		},
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
		{
			Input: []Range{
				{
					Lo: 2,
					Hi: 6,
				},
				{
					Lo: 4,
					Hi: 6,
				},
			},
			Output: true,
		},
		{
			Input: []Range{
				{
					Lo: 6,
					Hi: 8,
				},
				{
					Lo: 2,
					Hi: 4,
				},
			},
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := doRangesOverlapAtAll(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
