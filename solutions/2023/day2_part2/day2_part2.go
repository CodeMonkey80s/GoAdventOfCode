package day2_part2

import (
	"strconv"
	"strings"
)

func sumOfThePowerOfSets(games []string) int {
	var err error

	sum := 0

	scores := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, game := range games {
		sets := strings.Fields(game)
		colorName := ""
		colorValue := 0
		scores["red"] = 0
		scores["green"] = 0
		scores["blue"] = 0
		for i, word := range sets {
			if i == 0 {
				continue
			}
			if i == 1 {
				continue
			}
			if i%2 == 0 {
				colorValue, err = strconv.Atoi(strings.Trim(word, ":;,"))
				if err != nil {
					panic(err)
				}
			} else {
				colorName = strings.Trim(word, ":;,")
				if _, ok := scores[colorName]; ok {
					scores[colorName] = max(scores[colorName], colorValue)
				}
			}
			if i == len(sets)-1 {
				sum += scores["red"] * scores["green"] * scores["blue"]
			}
		}
	}

	return sum
}
