package day9_part1

import (
	"fmt"
	"reflect"
	"testing"

	"GoAdventOfCode/util"
)

func Test_moveHeadAndTail(t *testing.T) {
	var testCases = []struct {
		InputCommands []string
		OutputPoints  []point
		OutputVisited int
	}{
		{
			InputCommands: []string{
				"R 4",
			},
			OutputPoints: []point{
				{
					x: 4, y: 0,
				},
				{
					x: 3, y: 0,
				},
			},
			OutputVisited: 4,
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
			},
			OutputPoints: []point{
				{
					x: 4, y: -4,
				},
				{
					x: 4, y: -3,
				},
			},
			OutputVisited: 7,
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
			},
			OutputPoints: []point{
				{
					x: 1, y: -4,
				},
				{
					x: 2, y: -4,
				},
			},
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
			},
			OutputPoints: []point{
				{
					x: 1, y: -3,
				},
				{
					x: 2, y: -4,
				},
			},
			OutputVisited: 9,
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
			},
			OutputPoints: []point{
				{
					x: 5, y: -3,
				},
				{
					x: 4, y: -3,
				},
			},
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
			},
			OutputPoints: []point{
				{
					x: 5, y: -2,
				},
				{
					x: 4, y: -3,
				},
			},
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
			},
			OutputPoints: []point{
				{
					x: 0, y: -2,
				},
				{
					x: 1, y: -2,
				},
			},
		},
		{
			InputCommands: []string{
				"R 4",
				"U 4",
				"L 3",
				"D 1",
				"R 4",
				"D 1",
				"L 5",
				"R 2",
			},
			OutputPoints: []point{
				{
					x: 2, y: -2,
				},
				{
					x: 1, y: -2,
				},
			},
			OutputVisited: 13,
		},
	}
	lines := util.LoadInputFile("../inputs/day9_input.txt")
	testCase := []struct {
		InputCommands []string
		OutputPoints  []point
		OutputVisited int
	}{
		{
			InputCommands: lines,
			OutputVisited: 6391,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			ansHead, ansTail, ansVisited := getAnswer(tc.InputCommands)
			if len(tc.OutputPoints) > 0 {
				if !reflect.DeepEqual(ansHead, tc.OutputPoints[0]) {
					t.Errorf("Expected output of ansHead to be '%d' but we got '%d'", tc.OutputPoints[0], ansHead)
				}
				if !reflect.DeepEqual(ansTail, tc.OutputPoints[1]) {
					t.Errorf("Expected output of ansTail to be '%d' but we got '%d'", tc.OutputPoints[1], ansTail)
				}
			}
			if tc.OutputVisited > 0 {
				if ansVisited != tc.OutputVisited {
					t.Errorf("Expected output of ansVisited to be '%d' but we got '%d'", tc.OutputVisited, ansVisited)
				}
			}
		})
	}
}
