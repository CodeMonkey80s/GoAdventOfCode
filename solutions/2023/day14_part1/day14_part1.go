package day14_part1

import (
	"fmt"
)

func getAnswer(lines []string) int {

	grid := loadGrid(lines)

	// fmt.Println("BEFORE")
	// printGrid(grid)
	// fmt.Println()

	moveRocksN(grid)

	// fmt.Println("AFTER")
	// printGrid(grid)
	// fmt.Println()

	totalLoad := calculateTotalLoad(grid)

	return totalLoad
}

func moveRocksN(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'O' {
				ty := y - 1
				if ty < 0 {
					continue
				}
			fall:
				for ty >= 0 {
					switch {
					case grid[ty][x] == '#':
						break fall
					case grid[ty][x] == 'O':
						break fall
					default:
						grid[ty][x] = 'O'
						grid[ty+1][x] = '.'
						ty--
					}
				}
			}
		}
	}

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

func calculateTotalLoad(grid [][]rune) int {
	total := 0
	row := len(grid)
	for y := 0; y < len(grid); y++ {
		rocks := 0
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == 'O' {
				rocks++
			}
		}
		total += rocks * row
		row--
	}
	return total
}
