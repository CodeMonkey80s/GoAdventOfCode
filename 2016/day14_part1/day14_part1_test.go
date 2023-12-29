package day14_part1

import (
	"fmt"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputSalt string
		Output    int
	}{
		{
			InputSalt: "abc",
			Output:    22728,
		},
		{
			InputSalt: "yjdafjpo",
			Output:    25427,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputSalt)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
