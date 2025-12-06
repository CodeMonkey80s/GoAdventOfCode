package day3_part2

import (
	"GoAdventOfCode/util"
	"fmt"
	"testing"
)

var testCasesForPart2 = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"987654321111111"},
		Output: 987654321111,
	},
	{
		Input:  []string{"811111111111119"},
		Output: 811111111119,
	},
	{
		Input:  []string{"234234234234278"},
		Output: 434234234278,
	},
	{
		Input:  []string{"818181911112111"},
		Output: 888911112111,
	},
	{
		Input: []string{
			"987654321111111",
			"811111111111119",
			"234234234234278",
			"818181911112111",
		},
		Output: 3121910778619,
	},
}

func Test_getAnswerForPart1(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day3_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 167549941654721,
		},
	}
	testCasesForPart2 = append(testCasesForPart2, testCase...)
	for _, tc := range testCasesForPart2 {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswerForPart2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
