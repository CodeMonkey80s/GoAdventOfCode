package day13_part2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		// InputX marks starting point
		InputX int
		// InputY marks starting point y
		InputY int
		// InputW marks width of the map to be displayed
		InputW int
		// InputH marks height of the map to be displayed
		InputH int
		// InputN marks "office designer number"
		InputN int
		Output int
	}{
		{
			InputX: 7,
			InputY: 4,
			InputW: 50,
			InputH: 50,
			InputN: 10,
			Output: 151,
		},
		{
			InputX: 31,
			InputY: 39,
			InputW: 50,
			InputH: 60,
			InputN: 1364,
			Output: 127,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%vx%v_%v_%v\n", tc.InputX, tc.InputY, tc.InputN, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputN, tc.InputW, tc.InputH)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_makeOfficeBuilding(t *testing.T) {
	var testCases = []struct {
		InputW               int
		InputH               int
		Output               []string
		OutputNumberOfWalls  int
		OutputNumberOfFloors int
	}{
		{
			InputW:               10,
			InputH:               7,
			OutputNumberOfWalls:  33,
			OutputNumberOfFloors: 37,
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
			output, nw, nf := makeOfficeBuilding(tc.InputW, tc.InputH)
			if !reflect.DeepEqual(output, tc.Output) || nw != tc.OutputNumberOfWalls || nf != tc.OutputNumberOfFloors {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
