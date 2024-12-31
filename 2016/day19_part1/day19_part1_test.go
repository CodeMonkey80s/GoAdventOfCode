package day19_part1

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
			Input:  3005290,
			Output: 1816277,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getOutput(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.Output, output)
			}
		})
	}
}
