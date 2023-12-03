package day2_part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
		Output: 2286,
	},
}

func init() {
	f, err := os.Open("../inputs/day2_input.txt")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 65122,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_sumOfThePowerOfSets(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfThePowerOfSets(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
