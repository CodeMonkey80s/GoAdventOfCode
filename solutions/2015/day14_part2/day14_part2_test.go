package day14_part2

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input     []string
		InputTime int
		Output    int
	}{
		{
			Input: []string{
				"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds",
				"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds",
			},
			InputTime: 1000,
			Output:    689,
		},
	}
	lines := util.LoadInputFile("../inputs/day14_input.txt")
	testCase := []struct {
		Input     []string
		InputTime int
		Output    int
	}{
		{
			Input:     lines,
			InputTime: 2503,
			Output:    1256,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", "Example", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input, tc.InputTime)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getDeer(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output Deer
	}{
		{
			Input: "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds",
			Output: Deer{
				Name:       "Comet",
				Speed:      14,
				TimeAtMove: 10,
				TimeAtRest: 127,
				State:      stateFlying,
			},
		},
		{
			Input: "Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds",
			Output: Deer{
				Name:       "Dancer",
				Speed:      16,
				TimeAtMove: 11,
				TimeAtRest: 162,
				State:      stateFlying,
			},
		},
	}
	for _, tc := range testCases {
		parts := strings.Fields(tc.Input)
		name := parts[0]
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", name, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getDeer(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
