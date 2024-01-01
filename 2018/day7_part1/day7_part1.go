package day7_part1

import (
	"fmt"
	"slices"
	"strings"
)

func getAnswer(lines []string) string {

	E := make(map[byte][]byte)
	D := make(map[byte]int)
	Order := make([]byte, 0)
	for _, line := range lines {
		parts := strings.Fields(line)
		x := parts[1][0]
		y := parts[7][0]
		Order = append(Order, x)
		E[x] = append(E[x], y)
		D[y] += 1
	}

	Q := make([]byte, 0)
	for k, _ := range E {
		if D[k] == 0 {
			Q = append(Q, k)
		}
	}

	slices.Sort(Q)
	output := ""
	for len(Q) > 0 {
		x := Q[0]
		Q = Q[1:]
		output += string(x)
		for _, y := range E[x] {
			D[y] -= 1
			if D[y] == 0 {
				Q = append(Q, y)
			}
		}
		slices.Sort(Q)
	}
	fmt.Printf("Output: %s\n", output)

	return output
}
