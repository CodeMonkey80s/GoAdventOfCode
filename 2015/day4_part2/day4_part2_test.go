package day4_part2

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Name   string
	Input  string
	Output string
}{
	{
		Input:  "iwrupvqb",
		Output: "9958218",
	},
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
