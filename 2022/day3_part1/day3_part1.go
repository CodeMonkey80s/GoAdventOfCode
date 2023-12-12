package day3_part1

import (
	"log"
	"strings"
)

func calculateSum(lines []string) int {
	sum := 0
	for _, line := range lines {
		l := len(line)
		part1 := line[:l/2]
		part2 := line[l/2:]
		if len(part1) != len(part2) {
			log.Panicf("invalid part sizes, something is wrong with the input data?")
		}
		priorities := make(map[byte]int)
		for i := 0; i < len(part1); i++ {
			ch := part1[i]
			if idx := strings.IndexByte(part2, ch); idx != -1 {
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
