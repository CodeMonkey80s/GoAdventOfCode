package day13_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=10000000008400, Y=10000000005400",
				"",
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=10000000012748, Y=10000000012176",
				"",
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=10000000007870, Y=10000000006450",
				"",
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=10000000018641, Y=10000000010279",
				"",
			},
			Output: 0,
		},
	}
	lines := util.LoadInputFile("../inputs/day13_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 101214869433312,
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
