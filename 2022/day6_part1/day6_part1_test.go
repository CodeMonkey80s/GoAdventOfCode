package day6_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		Output: 7,
	},
	{
		Input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
		Output: 5,
	},
	{
		Input:  "nppdvjthqldpwncqszvftbrmjlhg",
		Output: 6,
	},
	{
		Input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		Output: 10,
	},
	{
		Input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		Output: 11,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day6_input.txt")
	testCase := []struct {
		Input  string
		Output int
	}{
		{
			Input:  lines[0],
			Output: 1896,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_startOfPacket(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := startOfPacket(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isMarker(t *testing.T) {
	testCases := []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "jpqm",
			Output: true,
		},
		{
			Input:  "mjqj",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isMarker(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
