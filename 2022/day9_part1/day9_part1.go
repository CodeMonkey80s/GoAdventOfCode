package day9_part1

import (
	"fmt"
	"math"
	"strings"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

var head point
var tail point

var visited = make(map[point]int)

func getAnswer(input []string) (point, point, int) {

	clear(visited)

	head = point{
		x: 0,
		y: 0,
	}
	tail = point{
		x: 0,
		y: 0,
	}

	for _, cmd := range input {
		head, tail = moveHeadAndTail(cmd, head, tail)
	}

	// fmt.Printf("visited: %d\n", len(visited))
	// printVisited()

	return head, tail, len(visited)
}

func moveHeadAndTail(cmd string, head point, tail point) (point, point) {

	parts := strings.Fields(cmd)
	dir := parts[0]
	length := util.ConvertStringToInt(parts[1])
	if length == 0 {
		panic("length is zero")
	}

	for i := 0; i < length; i++ {
		switch {
		case dir == "U":
			head.y--
		case dir == "D":
			head.y++
		case dir == "R":
			head.x++
		case dir == "L":
			head.x--
		default:
			panic("invalid direction")
		}

		// fmt.Printf("head: %+v\n", head)

		dist := int(math.Abs(float64(head.x-tail.x)) + math.Abs(float64(head.y-tail.y)))
		// fmt.Printf("dist: %d\n", dist)

		switch {
		case head.x == tail.x && head.y == tail.y:

		case dist >= 3 && head.y < tail.y && head.x > tail.x:
			tail.x++
			tail.y--
		case dist >= 3 && head.y > tail.y && head.x < tail.x:
			tail.x--
			tail.y++
		case dist >= 3 && head.y < tail.y && head.x < tail.x:
			tail.x--
			tail.y--
		case dist >= 3 && head.y > tail.y && head.x > tail.x:
			tail.x++
			tail.y++

		case dist >= 2 && head.x > tail.x && head.y == tail.y:
			tail.x++
		case dist >= 2 && head.x < tail.x && head.y == tail.y:
			tail.x--
		case dist >= 2 && head.y > tail.y && head.x == tail.x:
			tail.y++
		case dist >= 2 && head.y < tail.y && head.x == tail.x:
			tail.y--
		}

		visited[tail]++

		// fmt.Printf("tail: %+v\n", tail)
	}

	return head, tail
}

func printVisited() {

	x1 := 0
	x2 := 0

	y1 := 0
	y2 := 0

	for p := range visited {

		if p.x < x1 {
			x1 = p.x
		}

		if p.x > x2 {
			x2 = p.x
		}

		if p.y < y1 {
			y1 = p.y
		}

		if p.y > y2 {
			y2 = p.y
		}
	}

	for y := y1; y <= y2; y++ {
	loop:
		for x := x1; x <= x2; x++ {

			for p := range visited {
				if x == p.x && y == p.y {
					fmt.Print("#")
					continue loop
				}
			}
			fmt.Printf(".")
		}
		fmt.Println()
	}
}
