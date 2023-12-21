package day12_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  `[1,2,3]`,
		Output: 6,
	},
	{
		Input:  `[1,{"c":"red","b":2},3]`,
		Output: 4,
	},
	{
		Input:  `{"d":"red","e":[1,2,3,4],"f":5}`,
		Output: 0,
	},
	{
		Input:  `[1,"red",5]`,
		Output: 6,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Input  string
		Output int
	}{
		{
			Input:  lines[0],
			Output: 87842,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
