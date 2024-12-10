package day10_part2

import (
	"slices"

	"GoAdventOfCode/util"
)

type point struct {
	n int
	x int
	y int
}

func getAnswer(lines []string) int {
	topologyMap := loadTopologyMap(lines)

	score := 0
	for y := 0; y < len(topologyMap); y++ {
		for x := 0; x < len(topologyMap[0]); x++ {
			if topologyMap[y][x] == 0 {
				score += getTrailheadScore(topologyMap, x, y)
			}
		}
	}

	return score
}

func getTrailheadScore(topologyMap [][]int, x int, y int) int {
	p := point{
		n: topologyMap[y][x],
		x: x,
		y: y,
	}
	var points []point
	hash := make(map[point]int)
	points = append(points, p)
	points = traverseFrom(topologyMap, points, hash)
	// fmt.Printf("points: %v\n", points)
	// fmt.Printf("hash: %v\n", hash)

	score := 0
	for _, pts := range hash {
		score += pts
	}

	return score
}

func traverseFrom(topologyMap [][]int, points []point, hash map[point]int) []point {
	cx := points[len(points)-1].x
	cy := points[len(points)-1].y
	dx := 0
	dy := 0
	if points[len(points)-1].n == 9 {
		return points
	}
	if topologyMap[cy][cx] == '.' {
		return points
	}
	for dir := 0; dir < 4; dir++ {
		switch dir {
		case 0:
			dx = 0
			dy--
		case 1:
			dx++
			dy = 0
		case 2:
			dx = 0
			dy++
		default:
			dx--
			dy = 0
		}
		tx := cx + dx
		ty := cy + dy

		if tx < 0 || tx >= len(topologyMap[0]) || ty < 0 || ty >= len(topologyMap) {
			continue
		}

		if topologyMap[ty][tx] == '.' {
			continue
		}

		if topologyMap[cy][cx] == topologyMap[ty][tx]-1 {
			p := point{
				n: topologyMap[ty][tx],
				x: tx,
				y: ty,
			}
			points = append(points, p)
			points = traverseFrom(topologyMap, points, hash)

			if topologyMap[ty][tx] == 9 {
				hash[points[0]]++
				slices.Delete(points, len(points)-1, len(points))
				continue
			}

		}
	}

	slices.Delete(points, len(points)-1, len(points))

	return points
}

func loadTopologyMap(lines []string) [][]int {
	topologyMap := make([][]int, len(lines))
	for y, line := range lines {
		topologyMap[y] = make([]int, len(line))
		for x, ch := range line {
			if ch != '.' {
				n := util.ConvertRuneToInt(ch)
				topologyMap[y][x] = n
			} else {
				topologyMap[y][x] = -1
			}
		}
	}

	return topologyMap
}
