package day3_part2

import "fmt"

type Grid [][]rune

func howManyTrees(lines []string) int {

	grid, mx, my := NewGrid(lines)

	moves := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees := make([]int, 0)
	for _, delta := range moves {
		//fmt.Printf("Move %d,%d\n", delta[0], delta[1])
		n := move(grid, delta[0], delta[1])
		trees = append(trees, n)
	}

	//printGrid(grid)

	sum := 1
	for _, v := range trees {
		sum *= v
	}
	fmt.Printf("Size: %dx%d\n", mx, my)
	fmt.Printf("Trees: %d\n", sum)
	return sum
}

func NewGrid(lines []string) (Grid, int, int) {
	grid := make(Grid, 0)
	my := len(lines)
	mx := 0
	for i, line := range lines {
		grid = append(grid, make([]rune, len(line)))
		for j, ch := range line {
			grid[i][j] = ch
		}
		mx = max(mx, len(line))
	}
	return grid, mx, my
}

func move(grid Grid, dx int, dy int) int {
	x, y := 0, 0
	trees := 0
	for y < len(grid) {
		if grid[y][x] == '#' {
			//grid[y][x] = 'X'
			trees++
		} else {
			//grid[y][x] = 'O'
		}
		x = (x + dx) % len(grid[y])
		y += dy
	}
	return trees
}

/*

	https://en.wikipedia.org/wiki/ANSI_escape_code

	Black        0;30     Dark Gray     1;30
	Red          0;31     Light Red     1;31
	Green        0;32     Light Green   1;32
	Brown/Orange 0;33     Yellow        1;33
	Blue         0;34     Light Blue    1;34
	Purple       0;35     Light Purple  1;35
	Cyan         0;36     Light Cyan    1;36
	Light Gray   0;37     White         1;37

*/

func printGrid(grid Grid) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			switch {
			case grid[y][x] == 'O':
				fmt.Print("\033[37mO\033[0m")
			case grid[y][x] == '#':
				fmt.Printf("\033[1;32m%c\033[0m", grid[y][x])
			case grid[y][x] == 'X':
				fmt.Printf("\033[1;31m%c\033[0m", grid[y][x])
			default:
				fmt.Printf("\033[32m%c\033[0m", grid[y][x])
			}
		}
		fmt.Printf("\n")
	}
}
