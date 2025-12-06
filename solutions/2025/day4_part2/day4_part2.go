package day4_part2

import "fmt"

type point struct {
	X int
	Y int
}

func getAnswerForPart2(lines []string) int {

	grid := loadGrid(lines)

	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	isCountAsRoll := func(y int, x int, grid [][]rune) bool {
		if grid[y][x] != '@' {
			return false
		}

		var counter int
		for i := 0; i < len(directions); i++ {
			dy := directions[i][0]
			dx := directions[i][1]
			tx := x + dx
			ty := y + dy
			if tx < 0 || tx >= len(grid[0]) {
				continue
			}
			if ty < 0 || ty >= len(grid) {
				continue
			}
			if grid[ty][tx] == '@' {
				counter++
			}
		}

		if counter < 4 {
			return true
		}

		return false
	}

	var totalRolls int

	for {
		var (
			rolls  int
			points []point
		)

		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				ok := isCountAsRoll(y, x, grid)
				if ok {
					points = append(points, point{X: x, Y: y})
					rolls++
				}
			}
		}

		for _, p := range points {
			grid[p.Y][p.X] = '.'
		}

		//printGrid(grid)

		totalRolls += rolls

		if rolls == 0 {
			break
		}
	}

	return totalRolls
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
