package day3_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

type Area struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func getAnswer(lines []string) string {
	m := make(map[string]int)
	list := make(map[string]int)
	for i, line1 := range lines {
		id1, area1 := getArea(line1)
		list[id1] = 1
		for j, line2 := range lines {
			if j > i {
				id2, area2 := getArea(line2)
				ok := doAreasOverlap(area1, area2)
				//fmt.Printf("checking: %v and %v = %t\n", id1, id2, ok)
				if ok {
					m[id1]++
					m[id2]++
				}
			}
		}
	}
	for k := range list {
		_, ok := m[k]
		if !ok {
			return k
		}
	}

	return ""

}

func getArea(line string) (string, Area) {
	parts := strings.Fields(line)
	id := parts[0]
	p1 := strings.TrimSuffix(parts[2], ":")
	parts1 := strings.Split(p1, ",")
	parts2 := strings.Split(parts[3], "x")
	sx := util.ConvertStringToInt(parts1[0])
	sy := util.ConvertStringToInt(parts1[1])
	area := Area{
		X1: sx,
		Y1: sy,
		X2: sx + util.ConvertStringToInt(parts2[0]) - 1,
		Y2: sy + util.ConvertStringToInt(parts2[1]) - 1,
	}
	return id, area
}

func doAreasOverlap(a Area, b Area) bool {
	if a.X1 <= b.X2 && a.X2 >= b.X1 && a.Y1 <= b.Y2 && a.Y2 >= b.Y1 {
		return true
	}

	return false
}
