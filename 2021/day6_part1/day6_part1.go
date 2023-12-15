package day6_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

func howManyLanternfish(input string) int {
	list := make([]int, 0)
	parts := strings.Split(input, ",")
	for _, number := range parts {
		v := util.ConvertStringToInt(number)
		list = append(list, v)
	}
	day := 1
	after := 80
	for {
		for i, val := range list {
			list[i]--
			if val == 0 {
				list[i] = 6
				list = append(list, 8)
			}
		}
		//fmt.Printf("list: %d -> %v\n", day, IntToString(list))
		day++
		if day > after {
			break
		}
	}

	return len(list)
}

func IntToString(list []int) string {
	s := ""
	for _, v := range list {
		if len(s) > 0 {
			s += ","
		}
		s += util.ConvertIntToString(v)
	}

	return s
}
