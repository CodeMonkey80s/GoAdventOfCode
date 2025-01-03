package day12_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Label               string
		InputRegisters      map[string]int
		InputInstructions   []string
		OutputRegisterValue int
	}{
		{
			Label: "Instructions-sample",
			InputRegisters: map[string]int{
				"a": 0,
				"b": 0,
				"c": 0,
				"d": 0,
			},
			InputInstructions: []string{
				"cpy 41 a",
				"inc a",
				"inc a",
				"dec a",
				"jnz a 2",
				"dec a",
			},
			OutputRegisterValue: 42,
		},
	}
	lines := util.LoadInputFile("../inputs/day12_input.txt")
	testCase := []struct {
		Label               string
		InputRegisters      map[string]int
		InputInstructions   []string
		OutputRegisterValue int
	}{
		{
			Label: "Instructions-input",
			InputRegisters: map[string]int{
				"a": 0,
				"b": 0,
				"c": 0,
				"d": 0,
			},
			InputInstructions:   lines,
			OutputRegisterValue: 317993,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Label, tc.OutputRegisterValue)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputInstructions, tc.InputRegisters)
			if output != tc.OutputRegisterValue {
				t.Errorf("Expected output to be %v but we got %v", tc.OutputRegisterValue, output)
			}
		})
	}
}
