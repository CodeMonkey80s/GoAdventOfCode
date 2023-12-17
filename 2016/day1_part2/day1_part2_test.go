package day1_part2

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Name   string
	Input  string
	Output int
}{
	{
		Input:  "R8, R4, R4, R8",
		Output: 4,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day1_input.txt")
	testCase := []struct {
		Name   string
		Input  string
		Output int
	}{
		{
			Name:   "Puzzle Input",
			Input:  lines[0],
			Output: 153,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_getAnswer(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_Player_updateFacing(t *testing.T) {
	var testCases = []struct {
		Initial rune
		Input   rune
		Output  rune
	}{
		{
			Initial: 'N',
			Input:   'L',
			Output:  'W',
		},
		{
			Initial: 'N',
			Input:   'R',
			Output:  'E',
		},
		{
			Initial: 'S',
			Input:   'R',
			Output:  'W',
		},
		{
			Initial: 'S',
			Input:   'L',
			Output:  'E',
		},
		{
			Initial: 'E',
			Input:   'L',
			Output:  'N',
		},
		{
			Initial: 'E',
			Input:   'R',
			Output:  'S',
		},
		{
			Initial: 'W',
			Input:   'R',
			Output:  'N',
		},
		{
			Initial: 'W',
			Input:   'L',
			Output:  'S',
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %c %c Output: %c\n", tc.Initial, tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			p := Player{
				Facing: tc.Initial,
			}
			p.updateFacing(tc.Input)
			if p.Facing != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, p.Facing)
			}
		})
	}
}

func Test_getSteps(t *testing.T) {
	testCases := []struct {
		Input  string
		Output []Step
	}{
		{
			Input: "R81",
			Output: []Step{
				{
					Dir:    'R',
					Length: 81,
				},
			},
		},
		{
			Input: "L113",
			Output: []Step{
				{
					Dir:    'L',
					Length: 113,
				},
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getSteps(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
func Test_calculateDistance(t *testing.T) {
	testCases := []struct {
		InputA Point
		InputB Point
		Output int
	}{
		{
			InputA: Point{
				X: 0,
				Y: 0,
			},
			InputB: Point{
				X: 4,
				Y: 3,
			},
			Output: 7,
		},
		{
			InputA: Point{
				X: -2,
				Y: -2,
			},
			InputB: Point{
				X: 2,
				Y: 2,
			},
			Output: 8,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := calculateDistance(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
