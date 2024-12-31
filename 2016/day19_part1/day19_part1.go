package day19_part1

import (
	"fmt"
	"math"
)

func getOutput(value int) int {

	b := fmt.Sprintf("%b", uint(value))
	n := len(b) - 1
	l := value - int(math.Pow(2, float64(n)))

	return 2*l + 1
}
