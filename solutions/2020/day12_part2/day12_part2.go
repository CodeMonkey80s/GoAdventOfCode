package day12_part2

import (
	"fmt"
	"math"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

type waypointData struct {
	point
}

type shipData struct {
	point
}

func getAnswer(input []string) int {

	var waypoint waypointData
	var ship shipData
	ship.x = 0
	ship.y = 0
	waypoint.x = 10
	waypoint.y = -1

	moveShip := func(length int) {
		dx := waypoint.x - ship.x
		dy := waypoint.y - ship.y
		ship.x += dx * length
		ship.y += dy * length
		waypoint.x += dx * length
		waypoint.y += dy * length
	}

	moveWaypoint := func(dir string, length int) {
		switch dir {
		case "N":
			waypoint.y -= length
		case "S":
			waypoint.y += length
		case "E":
			waypoint.x += length
		case "W":
			waypoint.x -= length
		}
	}

	rotateWaypointL := func(p waypointData) (int, int) {
		adx := int(math.Abs(float64(waypoint.x - ship.x)))
		ady := int(math.Abs(float64(waypoint.y - ship.y)))
		dx := waypoint.x - ship.x
		dy := waypoint.y - ship.y
		switch {
		// E -> N
		case dx > 0 && dy <= 0:
			return ship.x - ady, ship.y - adx
		// N -> W
		case dx <= 0 && dy < 0:
			return ship.x - ady, ship.y + adx
		// W -> S
		case dx < 0 && dy >= 0:
			return ship.x + ady, ship.y + adx
		// S -> E
		case dx >= 0 && dy > 0:
			return ship.x + ady, ship.y - adx
		default:
			panic(fmt.Sprintf("something is wrong with rotation to left: %d, %d", dx, dy))
			return 0, 0
		}
	}

	rotateWaypointR := func(p waypointData) (int, int) {
		adx := int(math.Abs(float64(waypoint.x - ship.x)))
		ady := int(math.Abs(float64(waypoint.y - ship.y)))
		dx := waypoint.x - ship.x
		dy := waypoint.y - ship.y
		switch {
		// E -> S
		case dx > 0 && dy <= 0:
			return ship.x + ady, ship.y + adx
		// N -> E
		case dx <= 0 && dy < 0:
			return ship.x + ady, ship.y - adx
		// W -> N
		case dx < 0 && dy >= 0:
			return ship.x - ady, ship.y - adx
		// S -> W
		case dx >= 0 && dy > 0:
			return ship.x - ady, ship.y + adx
		default:
			panic(fmt.Sprintf("something is wrong with rotation to right: %d, %d", dx, dy))
			return 0, 0
		}
	}

	for _, line := range input {
		dir := line[:1]
		length := util.ConvertStringToInt(line[1:])

		switch dir {
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			moveWaypoint(dir, length)
		case "L":
			for i := 0; i < length/90; i++ {
				waypoint.x, waypoint.y = rotateWaypointL(waypoint)
			}
		case "R":
			for i := 0; i < length/90; i++ {
				waypoint.x, waypoint.y = rotateWaypointR(waypoint)
			}
		case "F":
			moveShip(length)
		}

		// dx := waypoint.x - ship.x
		// dy := waypoint.y - ship.y
		// fmt.Printf("ship: %d, %d, waypoint: %d, %d, diff: %d, %d\n", ship.x, ship.y, waypoint.x, waypoint.y, dx, dy)
		// fmt.Println()
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
