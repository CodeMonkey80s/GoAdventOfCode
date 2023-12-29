package day13_part1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputX int
		InputY int
		InputN int
		Output int
	}{
		{
			InputX: 7,
			InputY: 4,
			InputN: 10,
			Output: 11,
		},
		{
			InputX: 31,
			InputY: 39,
			InputN: 1364,
			Output: 86,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%vx%v_%v_%v\n", tc.InputX, tc.InputY, tc.InputN, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputX, tc.InputY, tc.InputN)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_makeOfficeBuilding(t *testing.T) {
	var testCases = []struct {
		InputW int
		InputH int
		Output []string
	}{
		{
			InputW: 10,
			InputH: 7,
			Output: []string{
				".#.####.##",
				"..#..#...#",
				"#....##...",
				"###.#.###.",
				".##..#..#.",
				"..##....#.",
				"#...##.###",
			},
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%vx%v_%v\n", tc.InputW, tc.InputH, tc.Output[0])
		t.Run(label, func(t *testing.T) {
			output := makeOfficeBuilding(tc.InputW, tc.InputH)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
