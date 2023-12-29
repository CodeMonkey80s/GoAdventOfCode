package day13_part2

import (
	"fmt"
	"math/bits"
)

type Point struct {
	X int
	Y int
}

func getAnswer(n int, w int, h int) int {

	sx := 1
	sy := 1
	steps := 0
	maxSteps := 50
	numberOfVisited := 0

	directions := []Point{
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: 0, Y: 1},
	}

	currentPos := make([]Point, 0)
	currentPos = append(currentPos, Point{
		X: sx,
		Y: sy,
	})
	visited := make(map[Point]bool)
	visited[Point{X: sx, Y: sy}] = true

	for steps < maxSteps {
		steps++
		nextPos := make([]Point, 0)
		for _, cp := range currentPos {
			for _, d := range directions {
				p := Point{
					X: cp.X + d.X,
					Y: cp.Y + d.Y,
				}
				_, ok := visited[p]
				if p.X >= 0 && p.Y >= 0 && !isWall(p.X, p.Y, n) && !ok {
					nextPos = append(nextPos, p)
					visited[p] = true
				}
			}

		}
		numberOfVisited = len(visited)
		currentPos = nextPos
	}

	//printLocations(visited, n, w, h)
	fmt.Printf("Locations: %d\n", numberOfVisited)
	return numberOfVisited
}

func printLocations(visited map[Point]bool, n int, w int, h int) {

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			_, ok := visited[Point{X: x, Y: y}]
			if ok {
				fmt.Print("\033[37mO\033[0m")
			} else if isWall(x, y, n) {
				fmt.Print("\033[31m#\033[0m")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func makeOfficeBuilding(w int, h int) ([]string, int, int) {

	numberOfWalls := 0
	numberOfFloors := 0
	building := make([][]rune, h)
	for y := 0; y < h; y++ {
		s := make([]rune, w)
		for x := 0; x < w; x++ {
			ok := isWall(x, y, w)
			if ok {
				s[x] = '#'
				numberOfWalls++
				continue
			}
			s[x] = '.'
			numberOfFloors++
		}
		building[y] = append(building[y], s...)
	}

	b := make([]string, 0, h)
	for _, floor := range building {
		b = append(b, string(floor))
	}
	return b, numberOfWalls, numberOfFloors
}

func isWall(x int, y int, n int) bool {
	i := uint((x*x + 3*x + 2*x*y + y + y*y) + n)
	c := bits.OnesCount(i)

	if c%2 == 1 {
		return true
	}

	return false
}
