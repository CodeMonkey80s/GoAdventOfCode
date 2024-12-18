package day8_part1

import (
	"fmt"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	grid := loadGrid(lines)
	printGrid(grid)

	visible := countVisible(grid)

	return visible
}

func countVisible(grid [][]rune) int {

	visible := (len(grid) * len(grid[0])) - ((len(grid) - 2) * (len(grid[0]) - 2))

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			visN := isVisible(grid, x, y, 0, -1)
			visS := isVisible(grid, x, y, 0, 1)
			visE := isVisible(grid, x, y, 1, 0)
			visW := isVisible(grid, x, y, -1, 0)
			if visN || visS || visW || visE {
				visible++
			}
		}
	}

	return visible
}

func isVisible(grid [][]rune, sx int, sy int, dx int, dy int) bool {

	ch := grid[sy][sx]
	sn := util.ConvertRuneToInt(ch)

	tx := sx
	ty := sy

	for {
		tx += dx
		ty += dy
		if tx < 0 || tx >= len(grid[0]) || ty < 0 || ty >= len(grid) {
			break
		}
		n := util.ConvertRuneToInt(grid[ty][tx])
		if n >= sn {
			return false
		}
	}

	return true
}

func loadGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		for x, ch := range line {
			grid[y][x] = ch
		}
	}

	return grid
}

func printGrid(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}
