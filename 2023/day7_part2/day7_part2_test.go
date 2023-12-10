package day7_part2

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"GoAdventOfCode/2023/util"
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
		Output: 5905,
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
			Output: 250825971,
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
			Input:  "4A4AA",
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

func Test_getCardsRankWithJoker(t *testing.T) {
	testCasesForCardRanks := []struct {
		Input  string
		Output int
	}{
		{
			Input:  "J213K",
			Output: cardRankOnePair,
		},
		{
			Input:  "1234J",
			Output: cardRankOnePair,
		},
		{
			Input:  "KKJ12",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "1134J",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "T55J5",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "QQQJA",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "KTJJT",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "QQJQQ",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "JJJJJ",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "7JJJ7",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "2QJJQ",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "JJ3J3",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "5999J",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "7Q7JJ",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "5JA6J",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "4444J",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "J1245",
			Output: cardRankOnePair,
		},
		{
			Input:  "J12J5",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "J1JJ2",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "JJJJ2",
			Output: cardRankFiveOfAKind,
		},
		{
			Input:  "J1J23",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "4JJ23",
			Output: cardRankThreeOfAKind,
		},
		{
			Input:  "JJJ23",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "QJJQ2",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "JKKK2",
			Output: cardRankFourOfAKind,
		},
		{
			Input:  "1122J",
			Output: cardRankFullHouse,
		},
		{
			Input:  "KKJ77",
			Output: cardRankFullHouse,
		},
	}
	for _, tc := range testCasesForCardRanks {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getCardsRankWithJoker(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_sortHands(t *testing.T) {
	input := make([]Hand, 0)
	for _, line := range testCases[0].Input {
		parts := strings.Fields(line)
		cards := parts[0]
		bid := util.ConvertStringToInt(parts[1])
		rank := 0
		if strings.ContainsRune(cards, Joker) {
			rank = getCardsRankWithJoker(cards)
		} else {
			rank = getCardsRank(cards)
		}
		h := Hand{
			Cards: cards,
			Bid:   bid,
			Rank:  rank,
		}
		input = append(input, h)
	}

	testCasesForSortingHands := []struct {
		Input  []Hand
		Output []Hand
	}{
		{
			Input: input,
			Output: []Hand{
				{
					Cards: "32T3K",
					Bid:   765,
					Rank:  2,
				},
				{
					Cards: "KK677",
					Bid:   28,
					Rank:  3,
				},
				{
					Cards: "T55J5",
					Bid:   684,
					Rank:  6,
				},
				{
					Cards: "QQQJA",
					Bid:   483,
					Rank:  6,
				},
				{
					Cards: "KTJJT",
					Bid:   220,
					Rank:  6,
				},
			},
		},
	}
	for _, tc := range testCasesForSortingHands {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sortHands(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
