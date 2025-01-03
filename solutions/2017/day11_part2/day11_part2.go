package day11_part2

import (
	"math"
	"strings"
)

type point struct {
	q int
	r int
	s int
}

func getAnswer(input string) int {

	var child point

	distance := 0

	steps := strings.Split(input, ",")
	for _, dir := range steps {
		switch dir {
		case "n":
			child.s++
			child.r--
		case "s":
			child.s--
			child.r++
		case "nw":
			child.q--
			child.s++
		case "ne":
			child.q++
			child.r--
		case "sw":
			child.q--
			child.r++
		case "se":
			child.q++
			child.s--
		}

		d := calculateDistance(point{}, child)
		distance = max(distance, d)
	}

	return distance
}

func calculateDistance(p1 point, p2 point) int {
	q := math.Abs(float64(p1.q - p2.q))
	r := math.Abs(float64(p1.q + p1.r - p2.q - p2.r))
	s := math.Abs(float64(p1.r - p2.r))
	return int(q+r+s) / 2
}
