package day7_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

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
			Output:   2797,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			_, a := getAnswer(tc.Input, tc.Register)
			clear(wires)
			wires["b"] = a
			_, output := getAnswer(tc.Input, tc.Register)
			if tc.Output != wires[tc.Register] {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
