package day7_part2

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getCombinations(t *testing.T) {

	var testCases = []struct {
		Input  int
		Output []string
	}{
		{
			Input: 3,
			Output: []string{
				"000",
				"001",
				"002",
				"010",
				"011",
				"012",
				"020",
				"021",
				"022",
				"100",
				"101",
				"102",
				"110",
				"111",
				"112",
				"120",
				"121",
				"122",
				"200",
				"201",
				"202",
				"210",
				"211",
				"212",
				"220",
				"221",
				"222",
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getCombinations([]rune{'0', '1', '2'}, tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"190: 10 19",
				"3267: 81 40 27",
				"83: 17 5",
				"156: 15 6",
				"7290: 6 8 6 15",
				"161011: 16 10 13",
				"192: 17 8 14",
				"21037: 9 7 18 13",
				"292: 11 6 16 20",
			},
			Output: 11387,
		},
	}
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 271691107779347,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
