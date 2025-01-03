package day3_part1

import (
	"fmt"
	"math"
	"strings"

	"GoAdventOfCode/util"
)

type Point struct {
	X int
	Y int
}

type Points map[Point]map[int]bool

func getAnswer(lines []string) int {

	points := make(Points)
	sp := Point{
		X: 0,
		Y: 0,
	}
	points[sp] = make(map[int]bool)
	points[sp][1] = true

	for i, line := range lines {
		cmds := strings.Split(line, ",")
		sx := sp.X
		sy := sp.Y
		n := i + 1
		for _, cmd := range cmds {
			//fmt.Printf("cmd: %s, sx: %d, sy: %d\n", cmd, sx, sy)
			points, sx, sy = doCommand(n, cmd, points, sx, sy)
		}
	}

	//printWires(points)

	distance := math.MaxUint32
	for p, v := range points {
		//fmt.Printf("Point: %v, V: %v\n", p, v)
		if len(v) == 2 {
			d := getManhattanDistance(p)
			//fmt.Printf("Point: %v Value: %v, Distance: %d\n", p, v, d)
			distance = min(distance, d)
		}
	}

	return distance
}

func getManhattanDistance(p2 Point) int {
	if p2.X < 0 {
		p2.X *= -1
	}
	if p2.Y < 0 {
		p2.Y *= -1
	}

	return p2.X + p2.Y
}

func doCommand(n int, cmd string, points Points, sx int, sy int) (Points, int, int) {
	dir := cmd[0]
	count := util.ConvertStringToInt(cmd[1:])
	ex := 0
	ey := 0
	for i := 1; i <= count; i++ {
		switch dir {
		case 'R':
			ex = sx + i
			ey = sy
		case 'L':
			ex = sx - i
			ey = sy
		case 'U':
			ex = sx
			ey = sy + i
		case 'D':
			ex = sx
			ey = sy - i
		}
		p := Point{X: ex, Y: ey}
		_, ok := points[p]
		if !ok {
			points[p] = make(map[int]bool)
		}
		points[p][n] = true
	}
	return points, ex, ey
}

func printWires(points Points) {

	sx := math.MaxInt32
	sy := math.MaxInt32
	mx := 0
	my := 0

	for p := range points {
		sx = min(sx, p.X)
		sy = min(sy, p.Y)
		mx = max(mx, p.X)
		my = max(my, p.Y)
	}

	for y := my; y >= sy; y-- {
		for x := sx; x <= mx; x++ {
			for p, v := range points {
				if p.X == x && p.Y == y {
					if x == 0 && y == 0 {
						fmt.Print("\033[32mO\033[0m")
					} else if len(v) == 1 {
						fmt.Print("\033[31m*\033[0m")
					} else if len(v) == 2 {
						fmt.Print("\033[33mX\033[0m")
					}
					goto loop
				}
			}
			fmt.Print("\033[34m.\033[0m")
		loop:
		}
		fmt.Println("")
	}
}
