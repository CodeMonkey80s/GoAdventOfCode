package day15_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswerForTwoIngredients(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			Output: 57600000,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswerForTwoIngredients(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day15_input.txt")
	testCases := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 117936,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswerForFourIngredients(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
