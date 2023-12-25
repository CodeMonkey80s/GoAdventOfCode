package day5_part2

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	{
		Input:  "abc",
		Output: "05ace8e3",
	},
	{
		Input:  "ojvtpuvg",
		Output: "1050cbbd",
	},
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getPassword(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
