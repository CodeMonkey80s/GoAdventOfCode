package day8_part2

import (
	"fmt"
	"testing"
)

// var testCases []struct {
//	Input  []string
//	Output int
// }

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"LR",
			"",
			"11A = (11B, XXX)",
			"11B = (XXX, 11Z)",
			"11Z = (11B, XXX)",
			"22A = (22B, XXX)",
			"22B = (22C, 22C)",
			"22C = (22Z, 22Z)",
			"22Z = (22B, 22B)",
			"XXX = (XXX, XXX)",
		},
		Output: 6,
	},
}

func init() {
	// lines := util.LoadInputFile("../inputs/day8_input.txt")
	// testCase := []struct {
	//	Input  []string
	//	Output int
	// }{
	//	{
	//		Input:  lines,
	//		Output: 28,
	//	},
	// }
	// testCases = append(testCases, testCase...)
}

func Test_howManySteps(t *testing.T) {
	t.SkipNow()
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input[0], tc.Output)
		t.Run(label, func(t *testing.T) {
			output := howManyStepsNodesEnd(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
