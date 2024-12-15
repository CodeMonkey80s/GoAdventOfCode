package day14_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func getAnswer(lines []string, cols int, rows int, turns int) int {

	var robots []robot

	safetyFactor := 0

	for _, line := range lines {
		s := strings.Replace(line, "p=", "", -1)
		s = strings.Replace(s, "v=", "", -1)
		s = strings.Replace(s, ",", " ", -1)
		parts := strings.Fields(s)
		r := robot{
			x:  util.ConvertStringToInt(parts[0]),
			y:  util.ConvertStringToInt(parts[1]),
			dx: util.ConvertStringToInt(parts[2]),
			dy: util.ConvertStringToInt(parts[3]),
		}
		robots = append(robots, r)
	}

	for turn := 1; turn <= turns; turn++ {
		for i := range robots {

			robots[i].x += robots[i].dx
			robots[i].y += robots[i].dy

			if robots[i].x < 0 {
				robots[i].x = robots[i].x%cols + cols
			}

			if robots[i].x >= cols {
				robots[i].x = robots[i].x % cols
			}

			if robots[i].y < 0 {
				robots[i].y = robots[i].y%rows + rows
			}

			if robots[i].y >= rows {
				robots[i].y = robots[i].y % rows
			}
		}

		// fmt.Printf("turn: %d\n", turn)
		// printRobots(robots, cols, rows)
	}

	q1 := countRobotsInQuadrant(robots, 0, (cols/2)-1, 0, (rows/2)-1)
	q2 := countRobotsInQuadrant(robots, (cols/2)+1, cols-1, 0, (rows/2)-1)
	q3 := countRobotsInQuadrant(robots, 0, (cols/2)-1, (rows/2)+1, rows-1)
	q4 := countRobotsInQuadrant(robots, (cols/2)+1, cols-1, (rows/2)+1, rows-1)
	// fmt.Printf("q1: %d\n", q1)
	// fmt.Printf("q2: %d\n", q2)
	// fmt.Printf("q3: %d\n", q3)
	// fmt.Printf("q4: %d\n", q4)

	safetyFactor = q1 * q2 * q3 * q4

	return safetyFactor
}

func countRobotsInQuadrant(robots []robot, x1 int, x2 int, y1 int, y2 int) int {
	count := 0
	for _, r := range robots {
		if r.x >= x1 && r.x <= x2 && r.y >= y1 && r.y <= y2 {
			count++
		}
	}
	return count
}

func printRobots(robots []robot, cols int, rows int) {
	type point struct {
		x int
		y int
	}
	cache := make(map[point]int)
	for _, r := range robots {
		p := point{
			x: r.x,
			y: r.y,
		}
		cache[p]++
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			ok := false
			for _, r := range robots {
				if !ok && r.x == x && r.y == y {
					p := point{
						x: x,
						y: y,
					}
					n, ok2 := cache[p]
					if ok2 {
						fmt.Printf(util.ConvertIntToString(n))
					}
					ok = true
				}
			}
			if !ok {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
