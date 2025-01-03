package day5_part2

import (
	"math"
	"slices"
	"strings"
)

func howManyUnits(input string) int {
	lowest := math.MaxUint32
	sb := strings.Builder{}
	for i := 'a'; i < 'z'; i++ {
		a := rune(i)
		b := rune(i - 32)
		for j := 0; j < len(input); j++ {
			if rune(input[j]) != a && rune(input[j]) != b {
				sb.WriteRune(rune(input[j]))
			}
		}
		polymers := []rune(sb.String())
		length := process(polymers)
		if length < lowest {
			lowest = length
		}
		sb.Reset()
	}

	return lowest
}

func process(polymers []rune) int {

	a := 0
	b := 1

loop:
	found := false
	for {
		if polymers[a]-32 == polymers[b] || polymers[a] == polymers[b]-32 {
			polymers = slices.Delete(polymers, a, b+1)
			found = true
			a = 0
			b = 1
			goto loop
		}
		a++
		b++
		if a >= len(polymers) || b >= len(polymers) {
			a = 0
			b = 1
			if !found {
				break
			}
			goto loop
		}
	}

	return len(polymers)
}
