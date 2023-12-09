package day7_part1

import (
	"slices"
	"strconv"
	"strings"
)

var cardsToValue = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

const (
	_ = iota
	FiveOfAKind
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

func totalWinnings(listOfHands []string) int {
	total := 0
	cardsToBid := make(map[string]int)
	listOfCards := make([]string, 0, len(listOfHands))
	for _, hand := range listOfHands {
		parts := strings.Fields(hand)
		cards := parts[0]
		bid := parts[1]
		cardsToBid[cards] = convertStringToInt(bid)
		listOfCards = append(listOfCards, cards)
	}

	slices.SortFunc(listOfCards, func(a, b string) int {
		va := getCardsRank(a)
		vb := getCardsRank(b)
		if va == vb {
			for i := 0; i < 5; i++ {
				if a[i] != b[i] {
					if cardsToValue[string(a[i])] < cardsToValue[string(b[i])] {
						return -1
					} else {
						return 1
					}
				}
			}
		}
		if va > vb {
			return -1
		}
		return 1
	})

	for i, card := range listOfCards {
		n := i + 1
		v, ok := cardsToBid[card]
		if ok {
			total += n * v
		}
	}

	return total
}

func getCardsRank(cards string) int {
	m := make(map[string]int)
	for _, ch := range cards {
		m[string(ch)]++
	}
	if len(m) == 1 {
		return FiveOfAKind
	}
	if len(m) == 5 {
		return HighCard
	}
	if len(m) == 4 {
		return OnePair
	}
	for _, v := range m {
		if v == 4 {
			return FourOfAKind
		}
		if v == 3 && len(m) == 3 {
			return ThreeOfAKind
		}
		if v == 2 && len(m) == 3 {
			return TwoPair
		}
	}

	return FullHouse
}

func convertStringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
