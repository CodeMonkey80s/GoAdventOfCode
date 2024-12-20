package day12_part2

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func getTotalPrice(lines []string) int {

	grid := loadGrid(lines)
	// printGrid(grid)

	isValid := func(points map[point]rune, x int, y int) bool {
		p := point{
			x: x,
			y: y,
		}
		_, ok := points[p]
		return ok
	}

	total := 0
	listOfPoints := loadPoints(grid)
	for _, points := range listOfPoints {
		area := len(points)
		corners := 0
		for p := range points {

			x := p.x
			y := p.y

			if !isValid(points, x-1, y) && !isValid(points, x, y-1) {
				corners++
			}
			if !isValid(points, x+1, y) && !isValid(points, x, y-1) {
				corners++
			}
			if !isValid(points, x-1, y) && !isValid(points, x, y+1) {
				corners++
			}
			if !isValid(points, x+1, y) && !isValid(points, x, y+1) {
				corners++
			}

			if isValid(points, x-1, y) && isValid(points, x, y-1) && !isValid(points, x-1, y-1) {
				corners++
			}
			if isValid(points, x+1, y) && isValid(points, x, y-1) && !isValid(points, x+1, y-1) {
				corners++
			}
			if isValid(points, x-1, y) && isValid(points, x, y+1) && !isValid(points, x-1, y+1) {
				corners++
			}
			if isValid(points, x+1, y) && isValid(points, x, y+1) && !isValid(points, x+1, y+1) {
				corners++
			}
		}

		total += area * corners
	}

	return total
}

func dfsPoints(grid [][]rune, points map[point]rune, ch rune, x int, y int) ([][]rune, map[point]rune) {

	if grid[y][x] == '.' || grid[y][x] != ch {
		return grid, points
	}

	p := point{
		x: x,
		y: y,
	}
	points[p] = ch
	grid[y][x] = '.'

	cols := len(grid[0])
	rows := len(grid)

	if x-1 >= 0 {
		grid, points = dfsPoints(grid, points, ch, x-1, y)
	}

	if y-1 >= 0 {
		grid, points = dfsPoints(grid, points, ch, x, y-1)
	}

	if x+1 < cols {
		grid, points = dfsPoints(grid, points, ch, x+1, y)
	}

	if y+1 < rows {
		grid, points = dfsPoints(grid, points, ch, x, y+1)
	}

	return grid, points
}

func loadPoints(grid [][]rune) []map[point]rune {

	listOfPoints := make([]map[point]rune, 0)

	cols := len(grid[0])
	rows := len(grid)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			ch := grid[y][x]
			if ch != '.' {
				points := make(map[point]rune)
				grid, points = dfsPoints(grid, points, ch, x, y)
				listOfPoints = append(listOfPoints, points)
			}
		}
	}

	return listOfPoints
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
