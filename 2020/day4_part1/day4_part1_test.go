package day4_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
				"byr:1937 iyr:2017 cid:147 hgt:183cm",
				"",
				"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
				"hcl:#cfa07d byr:1929",
				"",
				"hcl:#ae17e1 iyr:2013",
				"eyr:2024",
				"ecl:brn pid:760753108 byr:1931",
				"hgt:179cm",
				"",
				"hcl:#cfa07d eyr:2025 pid:166559648",
				"iyr:2011 ecl:brn hgt:59in",
			},
			Output: 2,
		},
	}
	lines := util.LoadInputFile("../inputs/day4_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 182,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			Output: true,
		},
		{
			Input:  "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			Output: false,
		},
		{
			Input:  "hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			Output: true,
		},
		{
			Input:  "hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", "Puzzle Input", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isValid(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
