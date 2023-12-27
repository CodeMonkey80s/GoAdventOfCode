package day4_part1

import (
	"fmt"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output int
	}{
		{
			Input:  "172851-675869",
			Output: -1,
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

func Test_verify(t *testing.T) {
	var testCases = []struct {
		Input  int
		Output bool
	}{
		{
			Input:  111111,
			Output: true,
		},
		{
			Input:  1234567,
			Output: false,
		},
		{
			Input:  12345,
			Output: false,
		},
		{
			Input:  223450,
			Output: false,
		},
		{
			Input:  123789,
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := verify(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
