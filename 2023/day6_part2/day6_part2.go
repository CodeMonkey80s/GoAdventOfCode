package day6_part2

import (
	"strconv"
	"strings"
)

func getTheNumber(lines []string) int {

	timeRaw := strings.TrimPrefix(lines[0], "Time:")
	timeRaw = strings.Replace(timeRaw, " ", "", -1)
	time := convertStringToInt(timeRaw)

	distanceRaw := strings.TrimPrefix(lines[1], "Distance:")
	distanceRaw = strings.Replace(distanceRaw, " ", "", -1)
	distance := convertStringToInt(distanceRaw)

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

func convertStringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
