package day11_part2

import (
	"fmt"
	"math"
)

const (
	charOccupied = '#'
	charEmpty    = 'L'
	charFloor    = '.'
)

func getOccupiedSeats(input []string, days int) int {

	grid := loadGrid(input)

	last := math.MaxInt
	for i := 1; i <= days; i++ {
		newGrid := copyGrid(grid)
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				ch := grid[y][x]
				switch {
				case ch == charEmpty:
					count := adjacent(grid, x, y, charOccupied)
					if count == 0 {
						newGrid[y][x] = charOccupied
					}
				case ch == charOccupied:
					count := adjacent(grid, x, y, charOccupied)
					if count >= 5 {
						newGrid[y][x] = charEmpty
					}
				}
			}
		}

		s := countSeats(newGrid, charOccupied)
		if s == last {
			break
		}

		grid = copyGrid(newGrid)
		// printGrid(grid)
		// fmt.Printf("day: %d, seats: %d\n", i, s)
	}

	seats := countSeats(grid, charOccupied)
	return seats
}

func adjacent(grid [][]rune, x int, y int, t rune) int {
	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
		{1, 1},
	}

	count := 0
outer:
	for _, d := range directions {
		i := 1
		for {
			tx := x + (i * d[0])
			ty := y + (i * d[1])
			if tx < 0 || ty < 0 || tx >= len(grid[0]) || ty >= len(grid) {
				continue outer
			}
			switch {
			case grid[ty][tx] == t:
				count++
				continue outer
			case grid[ty][tx] == charEmpty:
				continue outer
			}
			i++
		}
	}

	return count
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

func copyGrid(gridOrg [][]rune) [][]rune {
	grid := make([][]rune, len(gridOrg))
	for y := 0; y < len(gridOrg); y++ {
		grid[y] = make([]rune, len(gridOrg[0]))
		for x := 0; x < len(gridOrg[y]); x++ {
			grid[y][x] = gridOrg[y][x]
		}
	}

	return grid
}

func countSeats(grid [][]rune, t rune) int {
	seats := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == t {
				seats++
			}
		}
	}
	return seats
}

func printGrid(grid [][]rune) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}
