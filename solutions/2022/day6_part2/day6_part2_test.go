package day6_part2

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
		Output: 19,
	},
	{
		Input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
		Output: 23,
	},
	{
		Input:  "nppdvjthqldpwncqszvftbrmjlhg",
		Output: 23,
	},
	{
		Input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		Output: 29,
	},
	{
		Input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		Output: 26,
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
			Output: 3452,
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

func Test_isMessage(t *testing.T) {
	testCases := []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "tnvjfqwrcgsmlb",
			Output: true,
		},
		{
			Input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isMessage(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
