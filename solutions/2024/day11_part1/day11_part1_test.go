package day11_part1

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA  []int
		InputB  int
		OutputA []int
		OutputB int
	}{
		{
			InputA:  []int{0, 1, 10, 99, 999},
			InputB:  1,
			OutputA: []int{1, 2024, 1, 0, 9, 9, 2021976},
			OutputB: 7,
		},
		{
			InputA:  []int{125, 17},
			InputB:  6,
			OutputA: []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2},
			OutputB: 22,
		},
		{
			InputA:  []int{125, 17},
			InputB:  25,
			OutputA: []int{},
			OutputB: 55312,
		},
		{
			InputA:  []int{3028, 78, 973951, 5146801, 5, 0, 23533, 857},
			InputB:  25,
			OutputA: []int{},
			OutputB: 198089,
		},
		{
			InputA:  []int{3028, 78, 973951, 5146801, 5, 0, 23533, 857},
			InputB:  40,
			OutputA: []int{},
			OutputB: 0,
		},
	}
	for i, tc := range testCases {
		label := fmt.Sprintf("%v_%d\n", "Puzzle Input", i)
		t.Run(label, func(t *testing.T) {
			outputA, outputB := getAnswer(tc.InputA, tc.InputB)
			if len(tc.OutputA) > 0 {
				if !reflect.DeepEqual(outputA, tc.OutputA) {
					t.Errorf("Expected output to be of length '%d' but we got '%d'", tc.OutputA, outputA)
				}
			}
			if outputB != tc.OutputB {
				t.Errorf("Expected output to be of length '%d' but we got '%d'", tc.OutputB, outputB)
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
