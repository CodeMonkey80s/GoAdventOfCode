package day1_part1

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
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		},
		Output: 142,
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
			Output: 54990,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_sumOfAllTheCalibrationValues_part1(t *testing.T) {
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
