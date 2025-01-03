package day5_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

type Point struct {
	X int
	Y int
}

type Segment struct {
	P1 Point
	P2 Point
}

func howManyPointsOverlap(lines []string) int {
	mx := 0
	my := 0
	segments := make([]*Segment, 0)
	for _, line := range lines {
		p1, p2 := getPoints(line)
		if p1.X == p2.X || p1.Y == p2.Y {
			//fmt.Printf("i: %d, P1: %+v, p2: %+v\n", i, p1, p2)
			mx = max(mx, p1.X)
			mx = max(mx, p2.X)
			my = max(my, p1.Y)
			my = max(my, p2.Y)
			s := Segment{
				P1: p1,
				P2: p2,
			}
			segments = append(segments, &s)
		}
	}
	mx++
	my++

	diagram := createEmptyDiagram(mx, my)
	for _, seg := range segments {
		diagram = putSegmentOnDiagram(diagram, *seg)
	}

	//printDiagram(diagram)
	return countOverlaps(diagram)

}

func getPoints(line string) (Point, Point) {
	idx := strings.Index(line, "-")
	part1 := strings.TrimSpace(line[0:idx])
	part2 := strings.TrimSpace(line[idx+2:])
	parts1 := strings.Split(part1, ",")
	p1 := Point{
		X: util.ConvertStringToInt(parts1[0]),
		Y: util.ConvertStringToInt(parts1[1]),
	}
	parts2 := strings.Split(part2, ",")
	p2 := Point{
		X: util.ConvertStringToInt(parts2[0]),
		Y: util.ConvertStringToInt(parts2[1]),
	}
	return p1, p2
}

func createEmptyDiagram(mx int, my int) [][]int {
	diagram := make([][]int, 0)
	for y := 0; y < my; y++ {
		line := make([]int, mx)
		diagram = append(diagram, line)
	}
	return diagram
}

func printDiagram(diagram [][]int) {
	for y := 0; y < len(diagram); y++ {
		line := diagram[y]
		for x := 0; x < len(line); x++ {
			fmt.Print("", diagram[y][x])
		}
		fmt.Println()
	}
}

func putSegmentOnDiagram(diagram [][]int, seg Segment) [][]int {
	if seg.P1.X == seg.P2.X {
		if seg.P1.Y > seg.P2.Y {
			for y := seg.P2.Y; y <= seg.P1.Y; y++ {
				diagram[y][seg.P1.X]++
			}
		} else {
			for y := seg.P1.Y; y <= seg.P2.Y; y++ {
				diagram[y][seg.P1.X]++
			}
		}
	}
	if seg.P1.Y == seg.P2.Y {
		if seg.P1.X > seg.P2.X {
			for x := seg.P2.X; x <= seg.P1.X; x++ {
				diagram[seg.P1.Y][x]++
			}
		} else {
			for x := seg.P1.X; x <= seg.P2.X; x++ {
				diagram[seg.P1.Y][x]++
			}
		}
	}
	return diagram
}

func countOverlaps(diagram [][]int) int {
	overlap := 0
	for y := 0; y < len(diagram); y++ {
		line := diagram[y]
		for x := 0; x < len(line); x++ {
			if diagram[y][x] > 1 {
				overlap++
			}
		}
	}
	return overlap
}
