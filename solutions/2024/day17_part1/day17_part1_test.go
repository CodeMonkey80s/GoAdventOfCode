package day17_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input          []string
		OutputReg      rune
		OutputRegValue int
		OutputProgram  string
	}{
		{
			Input: []string{
				"Register A: 10",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 5,0,5,1,5,4",
			},
			OutputReg:      0,
			OutputRegValue: 0,
			OutputProgram:  "0,1,2",
		},
		{
			Input: []string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			OutputReg:      'a',
			OutputRegValue: 0,
			OutputProgram:  "4,2,5,6,7,7,7,7,3,1,0",
		},
		{
			Input: []string{
				"Register A: 0",
				"Register B: 29",
				"Register C: 0",
				"",
				"Program: 1,7",
			},
			OutputReg:      'b',
			OutputRegValue: 26,
			OutputProgram:  "",
		},
		{
			Input: []string{
				"Register A: 0",
				"Register B: 2024",
				"Register C: 43690",
				"",
				"Program: 4,0",
			},
			OutputReg:      'b',
			OutputRegValue: 44354,
		},

		{
			Input: []string{
				"Register A: 729",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			OutputReg:      0,
			OutputRegValue: 0,
			OutputProgram:  "4,6,3,5,6,3,5,2,1,0",
		},
	}
	lines := util.LoadInputFile("../inputs/day17_input.txt")
	testCase := []struct {
		Input          []string
		OutputReg      rune
		OutputRegValue int
		OutputProgram  string
	}{
		{
			Input:         lines,
			OutputProgram: "3,6,3,7,0,7,0,3,0",
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			registers, program := getOutput(tc.Input)
			if tc.OutputProgram != "" {
				if program != tc.OutputProgram {
					t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputProgram, program)
				}
			}
			if tc.OutputReg != 0 {
				if program != tc.OutputProgram {
					t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputProgram, program)
				}
				if registers[tc.OutputReg] != tc.OutputRegValue {
					t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputRegValue, registers[tc.OutputReg])
				}
			}
		})
	}
}
