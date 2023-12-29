package day14_part2

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
			Output:    22551,
		},
		{
			InputSalt: "yjdafjpo",
			Output:    22045,
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
