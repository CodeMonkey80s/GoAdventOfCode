package day20_part2

import (
	"fmt"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  int
		Output int
	}{
		{
			Input:  180,
			Output: 10,
		},
		{
			Input:  36_000_000,
			Output: 884520,
		},
	}
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
