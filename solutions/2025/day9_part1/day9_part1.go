package day9_part1

import (
	"GoAdventOfCode/util"
	"fmt"
	"strings"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}

func getAnswerForPart1(lines []string) int {

	var points []point

	for _, line := range lines {
		parts := strings.Split(line, ",")

		x := util.ConvertStringToInt(parts[0])
		y := util.ConvertStringToInt(parts[1])

		p := point{x, y}
		points = append(points, p)
	}

	var biggestArea int

	for idx1, p1 := range points {
		for idx2, p2 := range points {
			if idx1 == idx2 {
				continue
			}

			if p1.x == p2.x {
				continue
			}

			if p1.y == p2.y {
				continue
			}

			a := p2.x - p1.x
			if a < 0 {
				a = -a
			}
			a++

			b := p2.y - p1.y
			if b < 0 {
				b = -b
			}
			b++

			area := a * b
			//fmt.Printf("p1: %v, p2: %v, a: %v; b: %v, area: %v\n", p1, p2, a, b, area)
			biggestArea = max(biggestArea, area)
		}
	}

	return biggestArea
}
