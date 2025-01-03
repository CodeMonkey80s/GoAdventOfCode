package day4_part1

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
		Input:  "abcdef",
		Output: "609043",
	},
	{
		Input:  "pqrstuv",
		Output: "1048970",
	},
	{
		Input:  "iwrupvqb",
		Output: "346386",
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
