package day3_part1

import (
	"strconv"
)

func powerConsumption(diagnostic []string) int {
	vg := ""
	ve := ""
	for i := 0; i < len(diagnostic[0]); i++ {
		m := map[byte]int{
			'0': 0,
			'1': 0,
		}
		for j := 0; j < len(diagnostic); j++ {
			ch := diagnostic[j][i]
			m[ch]++
		}
		if m['0'] > m['1'] {
			vg += "0"
			ve += "1"
		} else {
			vg += "1"
			ve += "0"
		}
	}
	gamma, err := strconv.ParseInt(vg, 2, 64)
	if err != nil {
		panic(err)
	}
	epsilon, err := strconv.ParseInt(ve, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(gamma * epsilon)
}
