package day6_part1

import (
	"GoAdventOfCode/util"
	"fmt"
	"strings"
)

func getAnswerForPart1(lines []string) int {

	grid := loadGrid(lines)

	var sum int
	for x := 0; x < len(grid[0]); x++ {
		symbol := grid[len(grid)-1][x]
		//fmt.Printf("symbol: %v\n", symbol)

		colSum := util.ConvertStringToInt(grid[0][x])

		for y := 1; y < len(grid)-1; y++ {
			v := util.ConvertStringToInt(grid[y][x])
			//fmt.Printf("val: %d %d\n", colSum, v)
			switch {
			case symbol == "+":
				colSum += v
			case symbol == "*":
				colSum *= v
			}
		}

		//fmt.Printf("col: %d\n", colSum)

		sum += colSum
	}

	return sum
}

func loadGrid(lines []string) [][]string {
	grid := make([][]string, len(lines))
	for y := range lines {
		parts := strings.Fields(lines[y])
		grid[y] = make([]string, len(parts))
		for x, num := range parts {
			grid[y][x] = num
		}
	}

	return grid
}

func printGrid(grid [][]string) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fmt.Printf("%s ", grid[y][x])
		}
		fmt.Println()
	}
}
