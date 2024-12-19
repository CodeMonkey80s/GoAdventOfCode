package day22_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

type node struct {
	size  int
	used  int
	avail int
	use   int
}

var grid map[point]node

func getAnswer(input []string) int {

	grid = make(map[point]node)

	mx := 0
	my := 0

	for i, line := range input {
		if i == 0 || i == 1 {
			continue
		}

		parts := strings.Fields(line)

		coordStr := strings.ReplaceAll(parts[0], "/dev/grid/node-", "")
		coordStr = strings.ReplaceAll(coordStr, "x", "")
		coordStr = strings.ReplaceAll(coordStr, "y", "")

		coords := strings.Split(coordStr, "-")
		x := util.ConvertStringToInt(coords[0])
		y := util.ConvertStringToInt(coords[1])

		p := point{
			x: x,
			y: y,
		}

		mx = max(mx, x)
		my = max(my, y)

		n := node{
			size:  util.ConvertStringToInt(strings.TrimSuffix(parts[1], "T")),
			used:  util.ConvertStringToInt(strings.TrimSuffix(parts[2], "T")),
			avail: util.ConvertStringToInt(strings.TrimSuffix(parts[3], "T")),
			use:   util.ConvertStringToInt(strings.TrimSuffix(parts[4], "%")),
		}
		grid[p] = n
	}

	nodesCount := 0

	for p1, n1 := range grid {
		if n1.use > 0 {
			for p2, n2 := range grid {
				if p1 != p2 && n2.avail > n1.used {
					nodesCount++
				}
			}
		}
	}

	return nodesCount
}
