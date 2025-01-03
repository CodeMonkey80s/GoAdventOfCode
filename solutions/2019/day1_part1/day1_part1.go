package day1_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	sum := 0
	for _, line := range lines {
		sum += adjustMass(line)
	}
	return sum
}

func adjustMass(line string) int {
	mass := util.ConvertStringToInt(line)
	mass = mass / 3
	mass -= 2
	return mass
}
