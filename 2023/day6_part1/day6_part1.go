package day6_part1

import (
	"strings"

	"GoAdventOfCode/2023/util"
)

func getTheNumber(lines []string) int {
	sum := 1

	times := make([]int, 0)
	lineTimes := strings.TrimPrefix(lines[0], "Time:")
	timesRaw := strings.Fields(lineTimes)
	for _, num := range timesRaw {
		times = append(times, util.ConvertStringToInt(num))
	}

	distances := make([]int, 0)
	lineDistances := strings.TrimPrefix(lines[1], "Distance:")
	distancesRaw := strings.Fields(lineDistances)
	for _, num := range distancesRaw {
		distances = append(distances, util.ConvertStringToInt(num))
	}

	for n := range times {
		count := 0
		for i := 0; i <= times[n]; i++ {
			distance := countDistance(i, times[n])
			if distance > distances[n] {
				count++
			}
		}
		sum = sum * count
	}

	return sum
}

func countDistance(t1 int, t2 int) int {
	v := t1
	d := 0
	for i := 0; i < t2-t1; i++ {
		d += v
	}
	return d
}
