package day24_part2

import (
	"fmt"
	"maps"
)

type point struct {
	q int
	r int
	s int
}

const (
	colBlack = 'B'
	colWhite = 'W'
)

func getAnswer(input []string, days int) int {

	grid := make(map[point]rune)

	var p point

loop:
	for _, line := range input {

		p.s = 0
		p.q = 0
		p.r = 0

		i := 0
		for {
			switch {
			case line[i] == 'n' && line[i+1] == 'w':
				p = dir(p, "nw")
				i += 2
			case line[i] == 'n' && line[i+1] == 'e':
				p = dir(p, "ne")
				i += 2
			case line[i] == 's' && line[i+1] == 'e':
				p = dir(p, "se")
				i += 2
			case line[i] == 's' && line[i+1] == 'w':
				p = dir(p, "sw")
				i += 2
			case line[i] == 'e':
				p = dir(p, "e")
				i++
			case line[i] == 'w':
				p = dir(p, "w")
				i++
			default:
				panic("something is wrong!")
			}

			if i >= len(line) {
				grid = flip(grid, p)
				continue loop
			}
		}
	}

	for p, color := range grid {
		if color == colWhite {
			delete(grid, p)
		}
	}

	for i := 0; i < days; i++ {
		pointsCounter := make(map[point]int)
		for p := range grid {
			for _, d := range []string{"nw", "ne", "sw", "se", "w", "e"} {
				pointsCounter[dir(p, d)]++
			}
		}

		newGrid := make(map[point]rune)
		for p, count := range pointsCounter {
			if grid[p] == colBlack && count == 1 || count == 2 {
				newGrid[p] = colBlack
			}
		}

		clear(grid)
		maps.Copy(grid, newGrid)
	}

	tiles := 0
	for _, color := range grid {
		if color == colBlack {
			tiles++
		}
	}

	return tiles
}

func dir(p point, dir string) point {
	switch {
	case dir == "nw":
		p.r--
		p.s++
	case dir == "ne":
		p.r--
		p.q++
	case dir == "se":
		p.r++
		p.s--
	case dir == "sw":
		p.r++
		p.q--
	case dir == "e":
		p.q++
		p.s--
	case dir == "w":
		p.q--
		p.s++
	default:
		panic("something is wrong!")
	}

	return p
}

func flip(grid map[point]rune, p point) map[point]rune {
	color, ok := grid[p]
	if !ok {
		grid[p] = colBlack
		return grid
	}

	if color == colWhite {
		grid[p] = colBlack
		return grid
	}

	grid[p] = colWhite

	return grid
}

func colorsAround(grid map[point]rune, p point, color rune) int {
	count := 0
	for _, d := range []string{"nw", "ne", "sw", "se", "w", "e"} {
		c, ok := grid[dir(p, d)]
		if !ok {
			continue
		}
		if c == color {
			count++
		}
	}

	return count
}

func printGrid(grid map[point]rune) {
	cBlack := 0
	cWhite := 0
	for p, color := range grid {
		if color == colBlack {
			cBlack++
		}
		if color == colWhite {
			cWhite++
		}
		fmt.Printf("p: %v, %c\n", p, color)
	}
	fmt.Printf("black: %d, white: %d, total: %d\n", cBlack, cWhite, len(grid))
	fmt.Println()
}
