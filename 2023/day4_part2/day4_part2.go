package day4_part2

import (
	"strconv"
	"strings"
)

func howManyTotalScratchcards(cards []string) int {
	sum := 0

	scores := map[int]int{1: 1}

	getScore := func(line string) (int, int) {
		score := 0
		sides := strings.Split(line, "|")
		cardA := sides[0]
		cardB := sides[1]
		idx := strings.IndexByte(cardA, ':') + 1

		card := cardA[:idx]
		card = strings.TrimRight(card, ":")
		card = strings.TrimLeft(card, "Card ")
		id, err := strconv.Atoi(card)
		if err != nil {
			panic(err)
		}

		cardA = cardA[idx:]
		numbersA := strings.Fields(cardA)
		numbersB := strings.Fields(cardB)

		for _, b := range numbersB {
			for _, a := range numbersA {
				if a == b {
					score++
				}
			}
		}

		return id, score
	}

	for _, card := range cards {
		id, score := getScore(card)
		if _, ok := scores[id]; !ok {
			scores[id] = 1
		}

		if score == 0 {
			continue
		}

		for i := id + 1; i <= id+score; i++ {
			if _, ok := scores[i]; !ok {
				scores[i] = 1
			}
			scores[i] += scores[id]
		}
	}

	for _, v := range scores {
		sum += v
	}

	return sum
}
