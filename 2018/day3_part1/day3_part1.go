package day3_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

const (
	gridSize = 1000
)

type Area struct {
	SX int
	SY int
	EX int
	EY int
}

func getAnswer(lines []string) int {

	grid := createGrid(gridSize)
	for _, line := range lines {
		area := getArea(line)
		grid = claimArea(area, grid)
		//fmt.Printf("%v\n", area)
	}

	return calculateOverlappingClaims(grid)

}

func createGrid(s int) [][]int {
	grid := make([][]int, 0, s)
	for y := 0; y < s; y++ {
		row := make([]int, 0, s)
		for x := 0; x < s; x++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}

	return grid
}

func getArea(line string) Area {
	parts := strings.Fields(line)
	p1 := strings.TrimSuffix(parts[2], ":")
	parts1 := strings.Split(p1, ",")
	parts2 := strings.Split(parts[3], "x")
	sx := util.ConvertStringToInt(parts1[0])
	sy := util.ConvertStringToInt(parts1[1])
	area := Area{
		SX: sx,
		SY: sy,
		EX: sx + util.ConvertStringToInt(parts2[0]) - 1,
		EY: sy + util.ConvertStringToInt(parts2[1]) - 1,
	}
	return area
}

func claimArea(area Area, grid [][]int) [][]int {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if x >= area.SX && x <= area.EX &&
				y >= area.SY && y <= area.EY {
				grid[y][x]++
			}
		}
	}
	return grid
}

func calculateOverlappingClaims(grid [][]int) int {
	n := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x] >= 2 {
				n++
			}
		}
	}
	return n
}

func printGrid(grid [][]int) {
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%d", grid[y][x])
			}
		}
		fmt.Println()
	}
}
