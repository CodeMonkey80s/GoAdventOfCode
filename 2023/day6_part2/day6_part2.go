package day6_part2

import (
	"strings"

	"GoAdventOfCode/2023/util"
)

func getTheNumber(lines []string) int {

	timeRaw := strings.TrimPrefix(lines[0], "Time:")
	timeRaw = strings.Replace(timeRaw, " ", "", -1)
	time := util.ConvertStringToInt(timeRaw)

	distanceRaw := strings.TrimPrefix(lines[1], "Distance:")
	distanceRaw = strings.Replace(distanceRaw, " ", "", -1)
	distance := util.ConvertStringToInt(distanceRaw)

	a := 0
	b := time
	d := 0
	d1 := 0
	d2 := 0
	for i := 0; i <= time; i++ {
		d = i * (time - i)
		if d1 == 0 && d >= distance {
			d1 = a
		}
		d = i * (time - i)
		if d2 == 0 && d >= distance {
			d2 = b
		}
		if d1 != 0 && d2 != 0 {
			break
		}
		a++
		b--
	}
	return 1 + d2 - d1
}
