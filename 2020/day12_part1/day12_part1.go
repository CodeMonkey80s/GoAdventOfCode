package day12_part1

import (
	"fmt"
	"math"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

type shipData struct {
	point
	face string
}

func getAnswer(input []string) int {

	var ship shipData
	ship.x = 0
	ship.y = 0
	ship.face = "E"

	moveShip := func(dir string, length int) {
		dx := 0
		dy := 0
		switch dir {
		case "N":
			dy--
		case "S":
			dy++
		case "E":
			dx++
		case "W":
			dx--
		}
		for i := 0; i < length; i++ {
			ship.x += dx
			ship.y += dy
		}
	}

	changeFaceL := func(face string) string {
		switch face {
		case "N":
			return "W"
		case "W":
			return "S"
		case "S":
			return "E"
		case "E":
			return "N"
		default:
			return ""
		}
	}

	changeFaceR := func(face string) string {
		switch face {
		case "N":
			return "E"
		case "E":
			return "S"
		case "S":
			return "W"
		case "W":
			return "N"
		default:
			return ""
		}
	}

	for _, line := range input {
		dir := line[:1]
		length := util.ConvertStringToInt(line[1:])
		fmt.Printf("dir: %q, length: %d\n", dir, length)

		switch dir {
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			moveShip(dir, length)
		case "L":
			for i := 0; i < length/90; i++ {
				ship.face = changeFaceL(ship.face)
			}
		case "R":
			for i := 0; i < length/90; i++ {
				ship.face = changeFaceR(ship.face)
			}
		case "F":
			moveShip(ship.face, length)
		}

		fmt.Printf("coords: %d, %d\n", ship.x, ship.y)
	}

	p1 := point{}
	p2 := point{x: ship.x, y: ship.y}
	distance := calculateDistance(p1, p2)

	return distance
}

func calculateDistance(p1 point, p2 point) int {
	dx := int(math.Abs(float64(p1.x - p2.x)))
	dy := int(math.Abs(float64(p1.y - p2.y)))
	return dx + dy
}
