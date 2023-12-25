package day18_part1

import (
	"fmt"
)

type Grid struct {
	W     int
	H     int
	Cells [][]int
}

func NewGrid(w int, h int) *Grid {
	g := &Grid{
		W:     w,
		H:     h,
		Cells: make([][]int, 0),
	}
	for y := 0; y < h; y++ {
		g.Cells = append(g.Cells, make([]int, w))
	}
	return g
}

func getAnswer(lines []string, steps int) int {

	w := len(lines[0])
	h := len(lines)

	grid := NewGrid(w, h)

	grid.fill(lines)

	//grid.Print()
	for i := 0; i < steps; i++ {
		grid.play()
	}
	//grid.Print()

	return grid.countLightsOn()
}

func (g *Grid) fill(lines []string) {
	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				g.Cells[y][x] = 1
			}
		}
	}
}

func (g *Grid) getNeighbors(sy int, sx int) int {
	total := 0
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			rx := sx + x
			ry := sy + y
			if rx < 0 || ry < 0 || rx >= g.W || ry >= g.H || rx == sx && ry == sy {
				continue
			}
			if g.Cells[ry][rx] == 1 {
				total++
			}
		}
	}
	return total
}

func (g *Grid) play() {
	newCells := make([][]int, 0)
	for y := 0; y < g.H; y++ {
		newCells = append(newCells, make([]int, g.W))
	}
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			neighbors := g.getNeighbors(y, x)
			//fmt.Printf("%v,%v neighbors: %v\n", y, x, neighbors)
			switch {
			case g.Cells[y][x] == 1 && (neighbors == 2 || neighbors == 3):
				newCells[y][x] = 1
			case g.Cells[y][x] == 0 && neighbors == 3:
				newCells[y][x] = 1
			default:
				newCells[y][x] = 0
			}
		}
	}

	g.Cells = newCells
}

func (g *Grid) Print() {
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			if g.Cells[y][x] == 1 {
				fmt.Printf("\033[33m#\033[0m")
			} else {
				fmt.Printf("\033[37m.\033[0m")
			}
		}
		fmt.Printf("\n")
	}
}

func (g *Grid) countLightsOn() int {
	total := 0
	for y := 0; y < g.H; y++ {
		for x := 0; x < g.W; x++ {
			if g.Cells[y][x] == 1 {
				total++
			}
		}
	}
	return total
}
