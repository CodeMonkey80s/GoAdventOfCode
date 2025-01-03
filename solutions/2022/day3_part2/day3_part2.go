package day3_part2

import (
	"log"
	"strings"
)

func calculateSum(lines []string) int {

	sum := 0
	sacks := extractSacks(lines)
	for _, sack := range sacks {
		priorities := make(map[byte]int)
		for i := 0; i < len(sack[0]); i++ {
			ch := sack[0][i]
			idx1 := strings.IndexByte(sack[1], ch)
			idx2 := strings.IndexByte(sack[2], ch)
			if idx1 != -1 && idx2 != -1 {
				val := valueFromCharacter(ch)
				_, ok := priorities[ch]
				if !ok {
					priorities[ch] += val
				}
			}
		}
		for _, v := range priorities {
			sum += v
		}
	}
	return sum
}

func extractSacks(lines []string) [][]string {
	sacks := make([][]string, 0)
	sack := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		sack = append(sack, lines[i])
		if (i+1)%3 == 0 {
			sacks = append(sacks, sack)
			sack = sack[:0:0]
		}
	}
	return sacks
}

func valueFromCharacter(ch byte) int {
	val := 0
	switch {
	case ch >= 'a' && ch <= 'z':
		val = int(ch) - 96
	case ch >= 'A' && ch <= 'Z':
		val = int(ch) - 38
	default:
		log.Panicf("invalid character: %q", ch)
	}
	return val
}
