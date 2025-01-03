package day7_part2

import (
	"sort"
	"strings"

	"GoAdventOfCode/util"
)

const Joker = 'J'

var cardToValue = map[byte]int{
	'A':   13,
	'K':   12,
	'Q':   11,
	'T':   10,
	'9':   9,
	'8':   8,
	'7':   7,
	'6':   6,
	'5':   5,
	'4':   4,
	'3':   3,
	'2':   2,
	Joker: 1,
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

func getCardsRankWithJoker(cards string) int {
	m := make(map[rune]int)
	for _, char := range cards {
		m[char] += 1
	}
	maxElements := 0
	twoPairs := 0
	for _, v := range m {
		maxElements = max(maxElements, v)
		if v == 2 {
			twoPairs++
		}
	}
	if m[Joker] == 5 {
		return cardRankFiveOfAKind
	}
	if m[Joker] == 4 {
		return cardRankFiveOfAKind
	}
	if m[Joker] == 3 {
		if len(m) == 2 {
			return cardRankFiveOfAKind
		}
		return cardRankFourOfAKind
	}
	if m[Joker] == 2 {
		if len(m) == 4 && maxElements == 2 {
			return cardRankThreeOfAKind
		}
		if len(m) == 4 {
			return cardRankTwoPair
		}
		if len(m) == 3 && maxElements == 2 {
			return cardRankFourOfAKind
		}
		if len(m) == 3 && maxElements == 1 {
			return cardRankFullHouse
		}
		return cardRankFiveOfAKind
	}
	if m[Joker] == 1 {
		//fmt.Printf("CARD: %s, len: %d, max: %d\n", cards, len(m), maxElements)
		if len(m) == 4 && maxElements == 2 {
			return cardRankThreeOfAKind
		}
		if len(m) == 2 && maxElements == 4 {
			return cardRankFiveOfAKind
		}
		if len(m) == 3 && maxElements == 2 && twoPairs == 2 {
			return cardRankFullHouse
		}
		if len(m) == 3 && maxElements == 2 {
			return cardRankThreeOfAKind
		}
		if len(m) == 3 && maxElements == 3 {
			return cardRankFourOfAKind
		}

	}

	return cardRankOnePair
}
