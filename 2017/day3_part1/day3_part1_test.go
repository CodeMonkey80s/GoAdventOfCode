package day3_part1

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "1",
		Output: 0,
	},
	{
		Input:  "12",
		Output: 3,
	},
	{
		Input:  "23",
		Output: 2,
	},
	{
		Input:  "1024",
		Output: 31,
	},
	{
		Input:  "368078",
		Output: 371,
	},
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
