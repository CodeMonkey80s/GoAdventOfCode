package day2_part2

import (
	"log"
)

func calculateScore(lines []string) int {
	score := 0

	for _, line := range lines {

		// A = Rock
		// B = Paper
		// C = Scissors

		// X = Rock			1
		// Y = Paper		2
		// Z = Scissors		3
		enemy := line[0]
		player := line[2]

		s := roundScore(enemy, player)
		//fmt.Printf("p1: %c, p2: %c, score: %d\n", enemy, player, s)

		score += s
	}

	return score
}

func roundScore(enemy byte, player byte) int {
	score := 0

	// X = Lose
	// Y = Draw
	// Z = Win

	switch {

	// Win

	//  Scissors vs Rock
	case enemy == 'C' && player == 'Z':
		score += 1 + 6

	// Rock vs Paper
	case enemy == 'A' && player == 'Z':
		score += 2 + 6

	// Paper vs Scissors
	case enemy == 'B' && player == 'Z':
		score += 3 + 6

	// Loss

	// Rock vs Scissors
	case enemy == 'A' && player == 'X':
		score += 3 + 0

	// Paper vs Rock
	case enemy == 'B' && player == 'X':
		score += 1 + 0

	// Scissors vs Paper
	case enemy == 'C' && player == 'X':
		score += 2 + 0

	// Draw

	// Rock vs Rock
	case enemy == 'A' && player == 'Y':
		score += 1 + 3

	// Paper vs Paper
	case enemy == 'B' && player == 'Y':
		score += 2 + 3

	// Scissors vs Scissors
	case enemy == 'C' && player == 'Y':
		score += 3 + 3

	default:
		log.Panicf("%c = %c", enemy, player)
	}
	return score
}
