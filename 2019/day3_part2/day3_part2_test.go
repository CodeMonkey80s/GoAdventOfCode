package day3_part2

import (
	"fmt"
	"reflect"
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
				"R8,U5,L5,D3",
				"U7,R6,D4,L4",
			},
			Output: 30,
		},
		{
			Input: []string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			},
			Output: 610,
		},
		{
			Input: []string{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
				"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			},
			Output: 410,
		},
	}
	lines := util.LoadInputFile("../inputs/day3_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 101386,
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

func Test_doCommand(t *testing.T) {
	testCases := []struct {
		Input   string
		InputSX int
		InputSY int
		Output  Points
	}{
		{
			Input:   "R3",
			InputSX: 1,
			InputSY: 1,
			Output: Points{
				{X: 2, Y: 1}: {1: 1},
				{X: 3, Y: 1}: {1: 2},
				{X: 4, Y: 1}: {1: 3},
			},
		},
		{
			Input:   "L3",
			InputSX: 1,
			InputSY: 1,
			Output: Points{
				{X: 0, Y: 1}:  {1: 1},
				{X: -1, Y: 1}: {1: 2},
				{X: -2, Y: 1}: {1: 3},
			},
		},
		{
			Input:   "U3",
			InputSX: 1,
			InputSY: 1,
			Output: Points{
				{X: 1, Y: 2}: {1: 1},
				{X: 1, Y: 3}: {1: 2},
				{X: 1, Y: 4}: {1: 3},
			},
		},
		{
			Input:   "D3",
			InputSX: 1,
			InputSY: 1,
			Output: Points{
				{X: 1, Y: 0}:  {1: 1},
				{X: 1, Y: -1}: {1: 2},
				{X: 1, Y: -2}: {1: 3},
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			points := make(Points)
			points, _, _, _ = doCommand(1, 0, tc.Input, points, tc.InputSX, tc.InputSY)
			if !reflect.DeepEqual(points, tc.Output) {
				t.Errorf("Expected output to be %+v but we got %+v", tc.Output, points)
			}
		})
	}
}
