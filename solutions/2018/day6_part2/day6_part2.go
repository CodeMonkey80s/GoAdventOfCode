package day6_part2

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

type Point struct {
	Name string
	X    int
	Y    int
}

type Area struct {
	SX int
	SY int
	MX int
	MY int
}

type Grid [][]string

func getAnswer(lines []string, m int) int {

	points := NewPoints(lines)
	//printPoints(points)

	area := findArea(points)
	//fmt.Printf("Area: %+v\n", area)

	grid := createGrid(area, points)
	output := processPointsDistances(grid, points, m)
	fmt.Printf("Output: %d\n", output)

	return output
}

func NewPoints(lines []string) []Point {

	points := make([]Point, 0)

	for i, line := range lines {
		parts := strings.Split(line, ", ")
		x := util.ConvertStringToInt(parts[0])
		y := util.ConvertStringToInt(parts[1])
		p := Point{
			Name: util.ConvertIntToLetters(i + 1),
			X:    x,
			Y:    y,
		}
		points = append(points, p)
	}

	return points
}

func createGrid(area Area, points []Point) Grid {

	grid := make(Grid, area.MY)
	for y := 0; y < len(grid); y++ {
		grid[y] = make([]string, area.MX)
		for x := 0; x < len(grid[y]); x++ {
			grid[y][x] = "??"
		}
	}

	for _, p := range points {
		grid[p.Y][p.X] = p.Name
	}

	return grid
}

func getLargestArea(grid Grid) int {

	areas := make(map[string]int)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			ch := grid[y][x]

			v, ok := areas[ch]
			if ok && v == -1 {
				continue
			}

			if x == 0 || y == 0 || x == len(grid[y])-1 || y == len(grid)-1 {
				areas[ch] = -1
				continue
			}

			switch {
			case "AA" <= ch && ch <= "ZZ":
				ch = strings.ToLower(ch)
				areas[ch]++
			case "aa" <= ch && ch <= "zz":
				areas[ch]++
			}
		}
	}
	//util.PrintOrderedMap("Areas", areas)
	largest := 0
	for _, v := range areas {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func findArea(points []Point) Area {
	area := Area{
		SX: math.MaxUint32,
		SY: math.MaxUint32,
		MX: 0,
		MY: 0,
	}
	for _, p := range points {
		if p.X > area.MX {
			area.MX = p.X
		}
		if p.X < area.SX {
			area.SX = p.X
		}
		if p.Y > area.MY {
			area.MY = p.Y
		}
		if p.Y < area.SY {
			area.SY = p.Y
		}
	}
	area.SX--
	area.SY--
	area.MX += 2
	area.MY += 2
	return area
}

func processPointsDistances(grid Grid, points []Point, m int) int {
	size := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			dist := findLowestDistance(x, y, points)
			if dist < m {
				size++
			}
		}
	}

	return size
}

func findLowestDistance(x int, y int, points []Point) int {
	total := 0
	for _, p := range points {
		total += int(math.Abs(float64(x-p.X)) + math.Abs(float64(y-p.Y)))
	}

	return total
}

func printDistances(distances map[string]int) {
	keys := make([]string, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, v := range keys {
		fmt.Printf("%q %v\n", v, distances[v])
	}
}

func printPoints(points []Point) {
	for _, p := range points {
		fmt.Printf("%q (%d, %d)\n", p.Name, p.X, p.Y)
	}
}

/*

	Black        0;30     Dark Gray     1;30
	Red          0;31     Light Red     1;31
	Green        0;32     Light Green   1;32
	Brown/Orange 0;33     Yellow        1;33
	Blue         0;34     Light Blue    1;34
	Purple       0;35     Light Purple  1;35
	Cyan         0;36     Light Cyan    1;36
	Light Gray   0;37     White         1;37

*/

func printGrid(grid Grid) {
	//util.PrintSlice("grid", grid)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			s := grid[y][x]
			switch {
			case "AA" <= s && s <= "ZZ":
				fmt.Printf("\033[37m%s\033[0m", s)
			case "aa" <= s && s <= "zz":
				fmt.Printf("\033[32m%s\033[0m", s)
			case s == "..":
				fmt.Print("\033[1;30m..\033[0m")
			default:
				fmt.Print("\033[1;31m??\033[0m")
			}
		}
		fmt.Println()
	}
}
