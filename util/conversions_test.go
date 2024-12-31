package util

import (
	"fmt"
	"testing"
)

func Test_isNumber(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "-1",
			Output: true,
		},
		{
			Input:  "+1",
			Output: true,
		},
		{
			Input:  "a10",
			Output: false,
		},
		{
			Input:  "1a",
			Output: false,
		},
		{
			Input:  "#1",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Input: %v\n", tc.Input)
		t.Run(label, func(t *testing.T) {
			output := IsNumber(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.Output, output)
			}
		})
	}
}
