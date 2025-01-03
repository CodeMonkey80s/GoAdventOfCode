package day5_part2

import (
	"fmt"
	"sort"
)

func getAnswer(lines []string) int {

	var seatIDs []int
	for _, pass := range lines {
		row := findRow(pass)
		column := findColumn(pass)
		seatID := row*8 + column
		seatIDs = append(seatIDs, seatID)
	}

	missingSeatID := findMissingSeatID(seatIDs)

	fmt.Printf("Output: %d\n", missingSeatID)
	return missingSeatID
}

func findMissingSeatID(seatIDs []int) int {
	sort.Ints(seatIDs)

	for i := 1; i < len(seatIDs)-1; i++ {
		if seatIDs[i+1] != seatIDs[i]+1 {
			return seatIDs[i] + 1
		}
	}

	return -1 // Error: No missing seat found
}

func findRow(boardingPass string) int {
	lower := 0
	upper := 127

	for i := 0; i < 7; i++ {
		mid := (lower + upper) / 2
		if boardingPass[i] == 'F' {
			upper = mid
		} else if boardingPass[i] == 'B' {
			lower = mid + 1
		}
	}

	return lower
}

func findColumn(boardingPass string) int {
	lower := 0
	upper := 7

	for i := 7; i < 10; i++ {
		mid := (lower + upper) / 2
		if boardingPass[i] == 'L' {
			upper = mid
		} else if boardingPass[i] == 'R' {
			lower = mid + 1
		}
	}

	return lower
}
