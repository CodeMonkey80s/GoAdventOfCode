package day3_part2

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "1",
		Output: 1,
	},
	{
		Input:  "2",
		Output: 1,
	},
	{
		Input:  "3",
		Output: 2,
	},
	{
		Input:  "4",
		Output: 4,
	},
	{
		Input:  "5",
		Output: 5,
	},
	{
		Input:  "361527",
		Output: 0,
	},
}

func Test_getAnswer(t *testing.T) {
	t.SkipNow()
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
