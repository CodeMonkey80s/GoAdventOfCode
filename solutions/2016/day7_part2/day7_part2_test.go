package day7_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{}

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 242,
		},
	}
	testCases = append(testCases, testCase...)
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

func Test_isSSLSupported(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "aba[bab]xyz",
			Output: true,
		},
		{
			Input:  "xyx[xyx]xyx",
			Output: false,
		},
		{
			Input:  "aaa[kek]eke",
			Output: true,
		},
		{
			Input:  "zazbz[bzb]cdb",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			listSupernet, listHypernet := getNets(tc.Input)
			output := isSSLSupported(listSupernet, listHypernet)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_hasABA(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "aba",
			Output: true,
		},
		{
			Input:  "xyx",
			Output: true,
		},
		{
			Input:  "ioxoj",
			Output: true,
		},
		{
			Input:  "aaa",
			Output: false,
		},
		{
			Input:  "zaz",
			Output: true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			_, output := getABAs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_hasBAB(t *testing.T) {
	var testCases = []struct {
		Input    string
		InputABA string
		Output   bool
	}{
		{
			Input:    "abcbabxe",
			InputABA: "aba",
			Output:   true,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := hasBAB(tc.Input, tc.InputABA)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
