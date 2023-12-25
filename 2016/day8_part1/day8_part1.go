package day8_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

const (
	screenWidth  = 50
	screenHeight = 6
)

type Screen [][screenWidth]int

func getAnswer(lines []string) int {

	screen := make(Screen, screenHeight)
	for _, line := range lines {
		parts := strings.Fields(line)
		cmd := parts[0]

		if cmd == "rect" {
			numbers := strings.Split(parts[1], "x")
			a := util.ConvertStringToInt(numbers[0])
			b := util.ConvertStringToInt(numbers[1])
			screen = doCommandRect(screen, a, b)
			continue
		}

		if cmd == "rotate" {
			numbers := strings.Split(parts[2], "=")
			d := util.ConvertStringToInt(numbers[1])
			n := util.ConvertStringToInt(parts[4])
			if parts[1] == "column" {
				screen = doCommandRotateColumn(screen, d, n)
			} else {
				screen = doCommandRotateRow(screen, d, n)
			}
			continue
		}
	}

	printScreen(screen)

	return countVisiblePixels(screen)
}

func doCommandRect(screen Screen, a int, b int) Screen {
	for y := 0; y < b; y++ {
		for x := 0; x < a; x++ {
			screen[y][x] = 1
		}
	}
	return screen
}

func doCommandRotateColumn(screen Screen, x int, n int) Screen {
	for i := 0; i < n; i++ {
		temp := screen[screenHeight-1][x]
		for j := screenHeight - 1; j > 0; j-- {
			screen[j][x] = screen[j-1][x]
		}
		screen[0][x] = temp
	}
	return screen

}

func doCommandRotateRow(screen Screen, y int, n int) Screen {
	for i := 0; i < n; i++ {
		temp := screen[y][screenWidth-1]
		for j := screenWidth - 1; j > 0; j-- {
			screen[y][j] = screen[y][j-1]
		}
		screen[y][0] = temp
	}
	return screen
}

func countVisiblePixels(screen Screen) int {
	total := 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel == 1 {
				total++
			}
		}
	}
	return total
}

func printScreen(screen Screen) {
	for _, row := range screen {
		for _, pixel := range row {
			if pixel == 1 {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}
