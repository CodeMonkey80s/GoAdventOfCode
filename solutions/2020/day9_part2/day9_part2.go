package day9_part2

import (
	"math"

	"GoAdventOfCode/util"
)

func getAnswer(input []string, number int) int {

	size := 2
	n := 0

	for {
		numMin := math.MaxInt
		numMax := math.MinInt

		sum := 0
		for i := n; i < n+size; i++ {
			val := util.ConvertStringToInt(input[i])
			numMin = min(numMin, val)
			numMax = max(numMax, val)
			sum += val
		}
		if sum == number {
			return numMin + numMax
		}

		n++
		if n+size > len(input) {
			n = 0
			size++
		}

		if size-2 > len(input) {
			break
		}
	}

	return -1
}
