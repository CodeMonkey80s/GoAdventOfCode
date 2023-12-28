package day5_part1

import (
	"slices"
)

func howManyUnits(input string) int {

	polymers := []rune(input)
	polymers = process(polymers)

	return len(polymers)
}

func process(polymers []rune) []rune {

	a := 0
	b := 1

loop:
	found := false
	for {
		if polymers[a]-32 == polymers[b] || polymers[a] == polymers[b]-32 {
			//fmt.Printf("polymers (bef): %s\n", string(polymers))
			//fmt.Printf("Removing %q and %q\n", polymers[a], polymers[b])
			polymers = slices.Delete(polymers, a, b+1)
			//fmt.Printf("polymers (aft): %s\n", string(polymers))
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

	return polymers
}
