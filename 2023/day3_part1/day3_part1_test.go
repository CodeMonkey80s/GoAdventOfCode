package day3_part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"467..114..",
			"...*......",
			"..35..633.",
			"......#...",
			"617*......",
			".....+.58.",
			"..592.....",
			"......755.",
			"...$.*....",
			".664.598..",
		},
		Output: 4361,
	},
}

func init() {
	f, err := os.Open("../inputs/day3_input.txt")
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
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			lines = append(lines, text)
		}
	}
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 554003,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_sumOfAllOfThePartNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sumOfAllOfThePartNumbers(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
