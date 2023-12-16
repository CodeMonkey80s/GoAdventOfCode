package day1_part2

import (
	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	sum := 0
	for _, line := range lines {
		val := util.ConvertStringToInt(line)
		sum += adjustMass(val)
	}
	return sum
}

func adjustMass(mass int) int {
	fuel := 0
	for {
		mass = int(float64(mass)/3) - 2
		if mass <= 0 {
			break
		}
		fuel += mass
	}
	return fuel
}
