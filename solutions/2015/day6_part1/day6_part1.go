package day6_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type Light struct {
	On bool
}

type Coords struct {
	SX int
	SY int
	EX int
	EY int
}

const (
	gridSize = 1000
)

func getAnswer(lines []string) int {
	grid := createGrid(gridSize)
	for _, line := range lines {
		cmd, coords := getCommand(line)
		grid = doCommand(cmd, coords, grid)
	}

	return getNumberOfLights(grid)
}

func createGrid(s int) [][]*Light {
	grid := make([][]*Light, 0, s)
	for y := 0; y < s; y++ {
		row := make([]*Light, 0, s)
		for x := 0; x < s; x++ {
			light := Light{
				On: false,
			}
			row = append(row, &light)
		}
		grid = append(grid, row)
	}

	return grid
}

func getCommand(line string) (string, Coords) {
	idx := strings.IndexAny(line, "0123456789")
	parts := strings.Fields(line[idx:])
	cmd := strings.TrimSpace(line[:idx])
	cs := strings.Split(parts[0], ",")
	ce := strings.Split(parts[2], ",")
	coords := Coords{
		SX: util.ConvertStringToInt(cs[0]),
		SY: util.ConvertStringToInt(cs[1]),
		EX: util.ConvertStringToInt(ce[0]),
		EY: util.ConvertStringToInt(ce[1]),
	}
	return cmd, coords
}

func doCommand(cmd string, pos Coords, grid [][]*Light) [][]*Light {
	for y := pos.SY; y <= pos.EY; y++ {
		for x := pos.SX; x <= pos.EX; x++ {
			light := grid[y][x]
			switch cmd {
			case "toggle":
				light.On = !light.On
			case "turn off":
				light.On = false
			case "turn on":
				light.On = true
			}
			grid[y][x] = light
		}
	}

	return grid
}

func getNumberOfLights(grid [][]*Light) int {
	lights := 0
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			if grid[y][x].On {
				lights++
			}
		}
	}
	return lights
}
