package day15_part1

import (
	"fmt"
	"strings"
)

type robot struct {
	x int
	y int
}

var rob *robot

func getAnswer(lines []string) int {

	grid, moves := loadGrid(lines)

	// rows := len(grid)
	// cols := len(grid[0])
	//
	// fmt.Printf("grid: %d, %d\n", cols, rows)
	// fmt.Printf("string: %q\n", moves)
	// fmt.Printf("rob: %+v\n", rob)

	for _, move := range moves {

		switch move {
		case '^':
			moveRobot(grid, rob, 0, -1)
		case '<':
			moveRobot(grid, rob, -1, 0)
		case '>':
			moveRobot(grid, rob, 1, 0)
		case 'v':
			moveRobot(grid, rob, 0, 1)
		}

		// fmt.Printf("i: %d, %c\n", i, move)
	}

	// printGrid(grid, rob)

	sum := calculateCoordinateSum(grid)

	return sum
}

func isOnGrid(grid [][]rune, x int, y int) bool {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return false
	}
	return true
}

func moveRobot(grid [][]rune, rob *robot, dx int, dy int) {
	tx := rob.x + dx
	ty := rob.y + dy

	if !isOnGrid(grid, tx, ty) {
		return
	}

	if grid[ty][tx] == '#' {
		return
	}

	if grid[ty][tx] == '.' {
		rob.x = tx
		rob.y = ty
		return
	}

	if grid[ty][tx] == 'O' {
		ex := tx
		ey := ty
		for {
			ex += dx
			ey += dy
			if !isOnGrid(grid, ex, ey) {
				return
			}
			if grid[ey][ex] == '#' {
				return
			}
			if grid[ey][ex] == '.' {
				grid[ey][ex] = 'O'
				grid[ty][tx] = '.'
				rob.x = tx
				rob.y = ty
				return
			}
		}
	}
}

func loadGrid(lines []string) ([][]rune, string) {
	var moves string

	rows := 0
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		rows++
	}

	grid := make([][]rune, rows)
	for y, line := range lines {
		if len(line) == 0 {
			moves = strings.Join(lines[y+1:], "")
			break
		}
		grid[y] = make([]rune, len(line))
		for x, ch := range line {
			if ch == '@' {
				rob = new(robot)
				rob.x = x
				rob.y = y
				grid[y][x] = '.'
				continue
			}
			grid[y][x] = ch
		}
	}

	return grid, moves
}

func calculateCoordinateSum(grid [][]rune) int {
	sum := 0
	rows := len(grid[0])
	cols := len(grid)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 'O' {
				sum += y*100 + x
			}
		}
	}

	return sum
}

func printGrid(grid [][]rune, r *robot) {
	rows := len(grid[0])
	cols := len(grid)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if r.x == x && r.y == y {
				fmt.Print("@")
				continue
			}
			fmt.Printf("%c", grid[y][x])
		}
		fmt.Println()
	}
}
