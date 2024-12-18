package day8_part2

import (
	"fmt"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	grid := loadGrid(lines)
	// printGrid(grid)

	score := countTreesScore(grid)

	return score
}

func countTreesScore(grid [][]rune) int {

	score := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			scoreN := countScore(grid, x, y, 0, -1)
			scoreS := countScore(grid, x, y, 0, 1)
			scoreW := countScore(grid, x, y, -1, 0)
			scoreE := countScore(grid, x, y, 1, 0)

			total := scoreN * scoreS * scoreE * scoreW
			score = max(score, total)
		}
	}

	return score
}

func countScore(grid [][]rune, sx int, sy int, dx int, dy int) int {

	ch := grid[sy][sx]
	sn := util.ConvertRuneToInt(ch)

	tx := sx
	ty := sy

	score := 0

	for {
		tx += dx
		ty += dy
		if tx < 0 || tx >= len(grid[0]) || ty < 0 || ty >= len(grid) {
			break
		}
		n := util.ConvertRuneToInt(grid[ty][tx])
		if n >= sn {
			score++
			return score
		}
		score++
	}

	return score
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
