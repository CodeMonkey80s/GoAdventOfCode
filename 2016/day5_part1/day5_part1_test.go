package day5_part1

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
		Output: "18f47a30",
	},
	{
		Input:  "ojvtpuvg",
		Output: "4543c154",
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
