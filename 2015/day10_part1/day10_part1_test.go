package day10_part1

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "1113222113",
		Output: 252594,
	},
}

func Test_mutate(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
