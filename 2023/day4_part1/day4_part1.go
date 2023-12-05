package day3_part2

import (
	"strings"
)

func howManyPoints(cards []string) int {
	sum := 0
	winNumbers := make(map[int][]string)
	for cardNum, card := range cards {
		c := strings.Split(card, "|")
		cardA := c[0]
		cardB := c[1]
		idx := strings.IndexByte(cardA, ':') + 1
		cardA = cardA[idx:]
		numbersA := strings.Fields(cardA)
		numbersB := strings.Fields(cardB)
		winNumbers[cardNum] = append(winNumbers[cardNum], numbersA...)
		match := 0
		for _, number := range numbersB {
			for _, winNumber := range winNumbers[cardNum] {
				if winNumber == number {
					match++
				}
			}
		}
		total := 0
		if match > 0 {
			total = 1 << (match - 1)
		}
		sum += total
	}

	return sum
}
