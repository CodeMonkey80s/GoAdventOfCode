package day10_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: []string{
				"position=< 9,  1> velocity=< 0,  2>",
				"position=< 7,  0> velocity=<-1,  0>",
				"position=< 3, -2> velocity=<-1,  1>",
				"position=< 6, 10> velocity=<-2, -1>",
				"position=< 2, -4> velocity=< 2,  2>",
				"position=<-6, 10> velocity=< 2, -2>",
				"position=< 1,  8> velocity=< 1, -1>",
				"position=< 1,  7> velocity=< 1,  0>",
				"position=<-3, 11> velocity=< 1, -2>",
				"position=< 7,  6> velocity=<-1, -1>",
				"position=<-2,  3> velocity=< 1,  0>",
				"position=<-4,  3> velocity=< 2,  0>",
				"position=<10, -3> velocity=<-1,  1>",
				"position=< 5, 11> velocity=< 1, -2>",
				"position=< 4,  7> velocity=< 0, -1>",
				"position=< 8, -2> velocity=< 0,  1>",
				"position=<15,  0> velocity=<-2,  0>",
				"position=< 1,  6> velocity=< 1,  0>",
				"position=< 8,  9> velocity=< 0, -1>",
				"position=< 3,  3> velocity=<-1,  1>",
				"position=< 0,  5> velocity=< 0, -1>",
				"position=<-2,  2> velocity=< 2,  0>",
				"position=< 5, -2> velocity=< 1,  2>",
				"position=< 1,  4> velocity=< 2,  1>",
				"position=<-2,  7> velocity=< 2, -2>",
				"position=< 3,  6> velocity=<-1, -1>",
				"position=< 5,  0> velocity=< 1,  0>",
				"position=<-6,  0> velocity=< 2,  0>",
				"position=< 5,  9> velocity=< 1, -2>",
				"position=<14,  7> velocity=<-2,  0>",
				"position=<-3,  6> velocity=< 2, -1>",
			},
			InputB: 3,
			Output: 3,
		},
	}
	lines := util.LoadInputFile("../inputs/day10_input.txt")
	testCase := []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: lines,
			InputB: 10_101,
			Output: 10_101,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputA, tc.InputB)

			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
