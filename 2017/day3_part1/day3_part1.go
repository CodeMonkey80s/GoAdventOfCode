package day3_part1

import (
	"math"

	"GoAdventOfCode/util"
)

func getAnswer(input string) int {

	square := util.ConvertStringToInt(input)

	if square == 1 {
		return 0
	}

	// Find the ring or layer in which the square lies
	var ring int
	for i := 1; ; i += 2 {
		if square <= i*i {
			ring = (i - 1) / 2
			break
		}
	}

	// Calculate the Manhattan distance from the square to the nearest access port (square 1)
	// Find the side of the square
	sideLength := ring*2 + 1
	maxSquareInRing := sideLength * sideLength
	side := (maxSquareInRing - square) % (sideLength - 1)

	// Calculate the steps needed to reach the access port (square 1)
	steps := ring + int(math.Abs(float64(side-ring)))

	return steps
}
