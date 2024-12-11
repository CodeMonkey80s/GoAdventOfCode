package day11_part2

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA []int
		InputB int
		Output int
	}{
		{
			InputA: []int{0, 1, 10, 99, 999},
			InputB: 1,
			Output: 7,
		},
		{
			InputA: []int{125, 17},
			InputB: 6,
			Output: 22,
		},
		{
			InputA: []int{125, 17},
			InputB: 25,
			Output: 55312,
		},
		{
			InputA: []int{3028, 78, 973951, 5146801, 5, 0, 23533, 857},
			InputB: 25,
			Output: 198089,
		},
		{
			InputA: []int{3028, 78, 973951, 5146801, 5, 0, 23533, 857},
			InputB: 75,
			Output: 236302670835517,
		},
	}
	for i, tc := range testCases {
		label := fmt.Sprintf("%v_%d\n", "Puzzle Input", i)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be of length '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}

func Test_hasEvenDigits(t *testing.T) {
	var testCases = []struct {
		Input   int
		OutputA string
		OutputB bool
	}{
		{
			Input:   17,
			OutputA: "17",
			OutputB: true,
		},
		{
			Input:   2,
			OutputA: "",
			OutputB: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Input)
		t.Run(label, func(t *testing.T) {
			output, ok := hasEvenDigits(tc.Input)
			if ok != tc.OutputB {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputB, ok)
			}
			if !reflect.DeepEqual(output, tc.OutputA) {
				t.Errorf("Expected output to be '%v' but we got '%v'", tc.OutputA, output)
			}
		})
	}
}
