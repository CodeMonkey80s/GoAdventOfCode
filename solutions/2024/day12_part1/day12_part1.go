package day12_part1

import (
	"fmt"
	"slices"
)

type point struct {
	x int
	y int
}

func getTotalPrice(lines []string) int {

	grid := loadGrid(lines)
	// printGrid(grid)

	totalPrice := calculateTotalPriceOfGrid(grid)

	return totalPrice
}

var vis = make(map[point]bool)

func calculateTotalPriceOfGrid(grid [][]rune) int {
	clear(vis)

	totalPrice := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			p := point{
				x: x,
				y: y,
			}
			_, ok := vis[p]
			if ok {
				continue
			}

			area, perimeter := calculatePerimeterFor(grid, y, x)
			totalPrice += area * perimeter
		}
	}

	return totalPrice
}

func calculatePerimeterFor(grid [][]rune, sy int, sx int) (int, int) {

	dx := 0
	dy := 0
	area := 0
	perimeter := 0

	queue := make([]point, 0)

	p := point{
		x: sx,
		y: sy,
	}
	queue = append(queue, p)
	vis[p] = true

	for len(queue) > 0 {

		cx := queue[0].x
		cy := queue[0].y

		queue = queue[1:]
		per := 4
		area++

		for i := 0; i < 4; i++ {
			switch i {
			case 0:
				dx = 0
				dy = -1
			case 1:
				dx = 1
				dy = 0
			case 2:
				dx = 0
				dy = 1
			default:
				dx = -1
				dy = 0
			}
			tx := cx + dx
			ty := cy + dy

			if tx < 0 || tx >= len(grid[0]) || ty < 0 || ty >= len(grid) {
				continue
			}

			if grid[ty][tx] == grid[cy][cx] {
				per--
			}

			p := point{
				x: tx,
				y: ty,
			}
			_, ok := vis[p]
			if ok {
				continue
			}

			if grid[ty][tx] == grid[cy][cx] {
				queue = append([]point{p}, queue...)
				vis[p] = true
			}
		}

		perimeter += per
	}

	return area, perimeter
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

func printOrderedMap[T1 rune, T2 any](m map[T1]T2) {
	keys := make([]T1, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, v := range keys {
		fmt.Printf("%c => \"%+v\"\n", v, m[v])
	}
}
