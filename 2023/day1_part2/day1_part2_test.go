package day1_part2

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
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
		},
		Output: 281,
	},
}

func init() {
	f, err := os.Open("../inputs/day1_input.txt")
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
			Output: 54473,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_sumOfAllTheCalibrationValues(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfAllTheCalibrationValues(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
