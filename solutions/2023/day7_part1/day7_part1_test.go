package day7_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input: []string{
			"32T3K 765",
			"T55J5 684",
			"KK677 28",
			"KTJJT 220",
			"QQQJA 483",
		},
		Output: 6440,
	},
}

func init() {
	lines := util.LoadInputFile("../inputs/day7_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 251216224,
		},
	}
	testCases = append(testCases, testCase...)
}

func Test_totalWinnings(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", len(tc.Input), tc.Output)
		t.Run(label, func(t *testing.T) {
			output := totalWinnings(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_getCardsRank(t *testing.T) {
	testCasesForCardRanks := []struct {
		Input  string
		Output int
	}{
		{
			Input:  "11345",
			Output: cardRankOnePair,
		},
		{
			Input:  "32T3K",
			Output: cardRankOnePair,
		},
		{
			Input:  "KK677",
			Output: cardRankTwoPair,
		},
		{
			Input:  "34341",
			Output: cardRankTwoPair,
		},
		{
			Input:  "3Q373",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "34447",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "A3383",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "ATA3A",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "3QQQ3",
			Output: cardRankFullHouse,
		},
		{
			Input:  "3Q3Q3",
			Output: cardRankFullHouse,
		},
		{
			Input:  "44444",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "QQQQ2",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "34444",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "K2687",
			Output: cardRankHighCard,
		},
		{
			Input:  "12345",
			Output: cardRankHighCard,
		},
	}
	for _, tc := range testCasesForCardRanks {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getCardsRank(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
