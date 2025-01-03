package day10_part1

import (
	"GoAdventOfCode/util"
)

type point struct {
	n int
	x int
	y int
}

func getAnswer(lines []string) int {
	topologyMap := loadTopologyMap(lines)

	ends := make(map[point]int)

	score := 0
	for y := 0; y < len(topologyMap); y++ {
		for x := 0; x < len(topologyMap[0]); x++ {
			if topologyMap[y][x] == 0 {
				clear(ends)
				getTrailheadScore(topologyMap, x, y, ends)
				score += len(ends)
			}
		}
	}

	// util.PrintMap("ends", ends)
	// util.PrintSlice("topologyMap", topologyMap)

	return score
}

func getTrailheadScore(topologyMap [][]int, x int, y int, ends map[point]int) {
	if topologyMap[y][x] == '.' {
		return
	}
	p := point{
		n: topologyMap[y][x],
		x: x,
		y: y,
	}
	var path []point
	path = append(path, p)
	// fmt.Printf("starting from point %v\n", p)
	path = traverseFrom(topologyMap, path, ends)
}

func traverseFrom(topologyMap [][]int, path []point, ends map[point]int) []point {
	cx := path[len(path)-1].x
	cy := path[len(path)-1].y
	dx := 0
	dy := 0
	if topologyMap[cy][cx] == '.' {
		return path
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
		if tx >= 0 && tx < len(topologyMap[0]) && ty >= 0 && ty < len(topologyMap) && topologyMap[cy][cx] == topologyMap[ty][tx]-1 {
			if topologyMap[ty][tx] == '.' {
				continue
			}
			p := point{
				n: topologyMap[ty][tx],
				x: tx,
				y: ty,
			}
			path = append(path, p)
			if topologyMap[ty][tx] == 9 {
				ends[p]++
			}

			path = traverseFrom(topologyMap, path, ends)
		}
	}

	return path
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
