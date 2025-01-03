package day11_part1

import (
	"fmt"
	"testing"
)

func Test_isValidPassword(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "hijklmmn",
			Output: false,
		},
		{
			Input:  "abbceffg",
			Output: false,
		},
		{
			Input:  "abbcegjk",
			Output: false,
		},
		{
			Input:  "abcdffaa",
			Output: true,
		},
		{
			Input:  "ghjaabcc",
			Output: true,
		},
		{
			Input:  "abcdegab",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isPasswordValid(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_incrementPassword(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output string
	}{
		{
			Input:  "aaaaaaaa",
			Output: "aaaaaaab",
		},
		{
			Input:  "vzbxkghb",
			Output: "vzbxkghc",
		},
		{
			Input:  "vzbxkzhz",
			Output: "vzbxkzia",
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := incrementPassword(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output string
	}{
		{
			Input:  "abcdefgh",
			Output: "abcdffaa",
		},
		{
			Input:  "vzbxkghb",
			Output: "vzbxxyzz",
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
