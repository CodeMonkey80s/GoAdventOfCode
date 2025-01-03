package day19_part1

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"H => HO",
			"H => OH",
			"O => HH",
			"",
			"HOH",
		},
		Output: 4,
	},
	{
		Input: []string{
			"H => HO",
			"H => OH",
			"O => HH",
			"",
			"HOHOHO",
		},
		Output: 7,
	},
}

func Test_getAnswer(t *testing.T) {
	// lines := util.LoadInputFile("../inputs/day19_input.txt")
	// testCase := []struct {
	// 	Input  []string
	// 	Output int
	// }{
	// 	{
	// 		Input:  lines,
	// 		Output: 576,
	// 	},
	// }
	// testCases = append(testCases, testCase...)
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
