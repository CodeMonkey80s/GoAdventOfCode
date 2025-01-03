package day13_part1

import (
	"fmt"
	"math"
	"math/bits"
)

type Point struct {
	X     int
	Y     int
	Steps int
}

func getAnswer(ex int, ey int, n int) int {

	sx := 1
	sy := 1
	for {
		queue := []Point{
			{
				X: sx,
				Y: sy,
			},
		}
		visited := make(map[Point]bool)
		visited[Point{X: sx, Y: sy}] = true

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]

			if item.X == ex && item.Y == ey {
				steps := math.MaxUint32
				//util.PrintSlice("queue", queue)
				for _, v := range queue {
					if v.Steps < steps {
						steps = v.Steps
					}
				}
				fmt.Printf("Steps: %d\n", steps)
				return steps
			}

			neighbors := []Point{
				{X: item.X - 1, Y: item.Y},
				{X: item.X + 1, Y: item.Y},
				{X: item.X, Y: item.Y - 1},
				{X: item.X, Y: item.Y + 1},
			}

			for _, neighbor := range neighbors {
				if neighbor.X >= 0 && neighbor.Y >= 0 && !isWall(neighbor.X, neighbor.Y, n) && !visited[neighbor] {
					nb := neighbor
					nb.Steps = item.Steps + 1
					queue = append(queue, nb)
					visited[neighbor] = true
				}
			}
		}
	}

	return -1
}

func makeOfficeBuilding(w int, h int) []string {

	building := make([][]rune, h)
	for y := 0; y < h; y++ {
		s := make([]rune, w)
		for x := 0; x < w; x++ {
			ok := isWall(x, y, w)
			if ok {
				s[x] = '#'
				continue
			}
			s[x] = '.'
		}
		building[y] = append(building[y], s...)
	}

	b := make([]string, 0, h)
	for _, floor := range building {
		b = append(b, string(floor))
	}

	return b
}

func isWall(x int, y int, n int) bool {
	i := uint((x*x + 3*x + 2*x*y + y + y*y) + n)
	c := bits.OnesCount(i)

	if c%2 == 1 {
		return true
	}

	return false
}
