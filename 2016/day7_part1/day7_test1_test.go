package day7_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{}

func Test_getAnswer(t *testing.T) {
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 110,
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

func Test_isTLSSupported(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "abba[mnop]qrst",
			Output: true,
		},
		{
			Input:  "abcd[bddb]xyyx",
			Output: false,
		},
		{
			Input:  "aaaa[qwer]tyui",
			Output: false,
		},
		{
			Input:  "ioxxoj[asdfgh]zxcvbn",
			Output: true,
		},
		{
			Input:  "babzuaikmedruqsuuv[emlhynmvfhsigdryo]iyblsqlpplrlahtwr",
			Output: true,
		},
		{
			Input:  "pvjsstffnygbjafe[ztrjbalsthujslnitl]xjsoqghvrjncejwww[khwjgywxyglvhgz]kaxctpvhleqfmlm",
			Output: false,
		},
		{"abba[mnop]qrst", true},
		{"abcd[bddb]xyyx", false},
		{"aaaa[qwer]tyui", false},
		{"ioxxoj[asdfgh]zxcvbn", true},
		{"xyyx[xyyx]xyyx", false},
		{"abba[xyyx]xyyx", false},
		{"xyyx[abeo]owaxyyxnai", true},
		{"xyyx[xyyx]abba", false},
		{"xyyx[xyyx]xyyz", false},
		{"abcd[efgh]ijkl[mnop]qrst", false},
		{"abcd[efgh]ijkl[mnop]qrst[qwer]uvwx", false},
		{"abcd[efgh]ijkl[mnop]qrstqwer]uvwx", false},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isTLSSupported(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_isABBA(t *testing.T) {
	var testCases = []struct {
		Input  string
		Output bool
	}{
		{
			Input:  "abba",
			Output: true,
		},
		{
			Input:  "xyyx",
			Output: true,
		},
		{
			Input:  "ioxxoj",
			Output: true,
		},
		{
			Input:  "aaaa",
			Output: false,
		},
		{
			Input:  "xzaaaawc",
			Output: false,
		},
	}
	for _, tc := range testCases {
		label := fmt.Sprintf("%v_%v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isABBA(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
