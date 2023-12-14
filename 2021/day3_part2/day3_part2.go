package day3_part2

import (
	"slices"

	"GoAdventOfCode/util"
)

func supportRating(diagnostic []string) int {

	ratingsOxygen := make([]string, len(diagnostic))
	ratingsScrubber := make([]string, len(diagnostic))
	copy(ratingsOxygen, diagnostic)
	copy(ratingsScrubber, diagnostic)

	oxygenRating := getOxygenRatings(ratingsOxygen)
	//fmt.Printf("oxygen: %v\n", oxygenRating)

	scrubberRating := getScrubberRatings(ratingsScrubber)
	//fmt.Printf("scrubber: %v\n", scrubberRating)

	return oxygenRating * scrubberRating
}

func getOxygenRatings(ratings []string) int {
	pos := 0
	for {
		ch := getMostCommonValue(pos, ratings)
		ratings = slices.DeleteFunc(ratings, func(s string) bool {
			return s[pos] != ch
		})
		if len(ratings) == 1 {
			return util.ConvertBinaryStringToInt(ratings[0])
		}
		pos++
	}
}

func getMostCommonValue(pos int, ratings []string) byte {
	ch := '0'
	m := map[byte]int{
		'0': 0,
		'1': 0,
	}
	for i := 0; i < len(ratings); i++ {
		ch := ratings[i][pos]
		m[ch]++
	}
	switch {
	case m['0'] == m['1']:
		ch = '1'
	case m['0'] > m['1']:
		ch = '0'
	case m['0'] < m['1']:
		ch = '1'
	}
	return byte(ch)
}

func getScrubberRatings(ratings []string) int {
	pos := 0
	for {
		ch := getLeastCommonValue(pos, ratings)
		ratings = slices.DeleteFunc(ratings, func(s string) bool {
			return s[pos] != ch
		})
		if len(ratings) == 1 {
			return util.ConvertBinaryStringToInt(ratings[0])
		}
		pos++
	}
}

func getLeastCommonValue(pos int, ratings []string) byte {
	ch := '0'
	m := map[byte]int{
		'0': 0,
		'1': 0,
	}
	for i := 0; i < len(ratings); i++ {
		ch := ratings[i][pos]
		m[ch]++
	}
	switch {
	case m['0'] == m['1']:
		ch = '0'
	case m['0'] < m['1']:
		ch = '0'
	case m['0'] > m['1']:
		ch = '1'
	}
	return byte(ch)
}
