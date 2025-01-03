package day7_part1

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer_Example(t *testing.T) {
	var testCases = []struct {
		Input         []string
		InputRegister string
		Output        connections
	}{
		{
			Input: []string{
				"123 -> x",
				"456 -> y",
				"x AND y -> d",
				"x OR y -> e",
				"x LSHIFT 2 -> f",
				"y RSHIFT 2 -> g",
				"NOT x -> h",
				"NOT y -> i",
			},
			InputRegister: "h",
			Output: map[string]uint16{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			wires, output := getAnswer(tc.Input, tc.InputRegister)
			if !reflect.DeepEqual(wires, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer_Puzzle_Input(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCases := []struct {
		Input    []string
		Register string
		Output   uint16
	}{
		{
			Input:    lines,
			Register: "a",
			Output:   16076,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			_, output := getAnswer(tc.Input, tc.Register)
			if tc.Output != wires[tc.Register] {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
