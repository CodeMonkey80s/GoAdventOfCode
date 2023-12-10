package day7_part1

import (
	"sort"
	"strings"

	"GoAdventOfCode/util"
)

var cardToValue = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

const (
	cardRankFiveOfAKind  = 7
	cardRankFourOfAKind  = 6
	cardRankFullHouse    = 5
	cardRankThreeOfAKind = 4
	cardRankTwoPair      = 3
	cardRankOnePair      = 2
	cardRankHighCard     = 1
)

type Hand struct {
	Cards string
	Bid   int
	Rank  int
}

func totalWinnings(listOfHands []string) int {
	hands := make([]Hand, 0, len(listOfHands))
	for _, hand := range listOfHands {
		parts := strings.Fields(hand)
		cards := parts[0]
		bid := util.ConvertStringToInt(parts[1])
		h := Hand{
			Cards: cards,
			Bid:   bid,
			Rank:  getCardsRank(cards),
		}
		hands = append(hands, h)
	}

	hands = sortHands(hands)

	total := 0
	for i, hand := range hands {
		n := i + 1
		total += n * hand.Bid
	}

	return total
}

func sortHands(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Rank == hands[j].Rank {
			for k := 0; k < 5; k++ {
				if hands[i].Cards[k] == hands[j].Cards[k] {
					continue
				}
				card1 := hands[i].Cards[k]
				card2 := hands[j].Cards[k]
				return cardToValue[card1] < cardToValue[card2]
			}
		}
		return hands[i].Rank < hands[j].Rank
	})
	return hands
}

func getCardsRank(cards string) int {
	m := make(map[rune]int)
	for _, char := range cards {
		m[char]++
	}
	if len(m) == 1 {
		return cardRankFiveOfAKind
	}
	if len(m) == 5 {
		return cardRankHighCard
	}
	if len(m) == 4 {
		return cardRankOnePair
	}
	for _, v := range m {
		if v == 4 {
			return cardRankFourOfAKind
		}
		if v == 3 && len(m) == 3 {
			return cardRankThreeOfAKind
		}
		if v == 2 && len(m) == 3 {
			return cardRankTwoPair
		}
	}

	return cardRankFullHouse
}
