package day10_part1

import (
	"fmt"
	"math"
	"strings"

	"GoAdventOfCode/util"
)

type light struct {
	x  int
	y  int
	dx int
	dy int
}

func getAnswer(lines []string, turns int) int {

	var lights []light

	for _, line := range lines {
		s := strings.Replace(line, "position=<", "", -1)
		s = strings.Replace(s, "velocity=<", "", -1)
		s = strings.Replace(s, ">", "", -1)
		s = strings.Replace(s, ",", " ", -1)
		parts := strings.Fields(s)
		l := light{
			x:  util.ConvertStringToInt(parts[0]),
			y:  util.ConvertStringToInt(parts[1]),
			dx: util.ConvertStringToInt(parts[2]),
			dy: util.ConvertStringToInt(parts[3]),
		}
		// fmt.Printf("p: %+v\n", l)
		lights = append(lights, l)
	}

	areas := make(map[int]int)

	for turn := 1; turn <= turns; turn++ {

		x1 := 0
		x2 := 0
		y1 := 0
		y2 := 0

		for i := range lights {

			lights[i].x += lights[i].dx
			lights[i].y += lights[i].dy

			if lights[i].x < x1 {
				x1 = lights[i].x
			}

			if lights[i].x > x2 {
				x2 = lights[i].x
			}

			if lights[i].y < y1 {
				y1 = lights[i].y
			}

			if lights[i].y > y2 {
				y2 = lights[i].y
			}

			area := int(math.Abs(float64(x2-x1)) * math.Abs(float64(y2-y1)))
			areas[turn] = area
		}

	}

	printLights(lights)

	turn := 0
	lowest := math.MaxInt
	for t, area := range areas {
		if area < lowest {
			lowest = area
			turn = t
		}
	}
	return turn
}

func printLights(lights []light) {

	x1 := 0
	x2 := 0

	y1 := 0
	y2 := 0

	for i := range lights {

		if lights[i].x < x1 {
			x1 = lights[i].x
		}

		if lights[i].x > x2 {
			x2 = lights[i].x
		}

		if lights[i].y < y1 {
			y1 = lights[i].y
		}

		if lights[i].y > y2 {
			y2 = lights[i].y
		}
	}

	for y := y1; y <= y2; y++ {
	loop:
		for x := x1; x <= x2; x++ {

			for i := range lights {
				if x == lights[i].x && y == lights[i].y {
					fmt.Print("#")
					continue loop
				}
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
}
