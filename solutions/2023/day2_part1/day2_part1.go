package day2_part1

import (
	"strconv"
	"strings"
)

func sumOfTheIDsOfGames(games []string) int {
	var err error

	const maxScoreRed int = 12
	const maxScoreGreen int = 13
	const maxScoreBlue int = 14

	sum := 0

	scores := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
next:
	for _, game := range games {
		sets := strings.Fields(game)
		gameID := 0
		colorName := ""
		colorValue := 0
		for i, word := range sets {
			if i == 0 {
				continue
			}
			if i == 1 {
				gameID, err = strconv.Atoi(strings.Trim(word, ":;,"))
				if err != nil {
					panic(err)
				}
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
					scores[colorName] += colorValue
				}
			}
			if strings.HasSuffix(word, ";") || i == len(sets)-1 {
				if scores["red"] > maxScoreRed ||
					scores["green"] > maxScoreGreen ||
					scores["blue"] > maxScoreBlue {
					scores["red"] = 0
					scores["green"] = 0
					scores["blue"] = 0
					continue next
				}
				scores["red"] = 0
				scores["green"] = 0
				scores["blue"] = 0
			}
		}
		sum += gameID
	}

	return sum
}
