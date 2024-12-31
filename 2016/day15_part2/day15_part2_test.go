package day15_part2

import (
	"fmt"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"Disc #1 has 17 positions; at time=0, it is at position 5.",
				"Disc #2 has 19 positions; at time=0, it is at position 8.",
				"Disc #3 has 7 positions; at time=0, it is at position 1.",
				"Disc #4 has 13 positions; at time=0, it is at position 7.",
				"Disc #5 has 5 positions; at time=0, it is at position 1.",
				"Disc #6 has 3 positions; at time=0, it is at position 0.",
				"Disc #7 has 11 positions; at time=0, it is at position 0.",
			},
			Output: 3543984,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Input)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
