package day6_part1

import "fmt"

type field struct {
	x     int
	y     int
	block bool
}

type coord struct {
	x int
	y int
}

var visited map[coord]bool

type guardData struct {
	x      int
	y      int
	facing rune
}

var guard guardData

func getAnswer(lines []string) int {

	visited = make(map[coord]bool)

	room := loadRoom(lines)

	guard.facing = 'N'

	guardWalk(room)
	// printRoom(room)

	return len(visited)
}

func getDirection(facing rune) (int, int) {
	switch facing {
	case 'N':
		return 0, -1
	case 'S':
		return 0, 1
	case 'W':
		return -1, 0
	case 'E':
		return 1, 0
	}

	return 0, 0
}

func guardWalk(room [][]field) {
	c := coord{
		x: guard.x,
		y: guard.y,
	}
	visited[c] = true
	for {
		dx, dy := getDirection(guard.facing)
		if dy+guard.y < 0 || dx+guard.x < 0 || dy+guard.y > len(room)-1 || dx+guard.x > len(room[0])-1 {
			return
		}
		switch room[guard.y+dy][guard.x+dx].block {
		case true:
			switch guard.facing {
			case 'N':
				guard.facing = 'E'
			case 'S':
				guard.facing = 'W'
			case 'W':
				guard.facing = 'N'
			case 'E':
				guard.facing = 'S'
			}
		default:
			guard.x += dx
			guard.y += dy
			c := coord{
				x: guard.x,
				y: guard.y,
			}
			visited[c] = true
		}
	}
}

func loadRoom(lines []string) [][]field {
	room := make([][]field, len(lines))
	for y, line := range lines {
		room[y] = make([]field, len(line))
		for x, ch := range line {
			f := field{
				x:     x,
				y:     y,
				block: false,
			}
			switch ch {
			case '^':
				guard.x = x
				guard.y = y
			case '#':
				f.x = x
				f.y = y
				f.block = true
				room[y][x] = f
			default:
				room[y][x] = f
			}
		}
	}

	return room
}

func printRoom(room [][]field) {
	for y := 0; y < len(room); y++ {
		for x := 0; x < len(room[y]); x++ {
			c := coord{
				x: x,
				y: y,
			}
			switch {
			case visited[c]:
				fmt.Print("X")
			case room[y][x].block:
				fmt.Print("#")
			case y == guard.y && x == guard.x:
				fmt.Print("@")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
