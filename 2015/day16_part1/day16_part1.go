package day16_part1

import (
	"strings"

	"GoAdventOfCode/util"
)

type Mapping = map[string]int

func getAnswer(lines []string) int {

	m := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for _, line := range lines {
		id, mapping := getMappingFromLine(line)
		found := true
		for k, v1 := range mapping {
			if m[k] != v1 {
				found = false
			}

		}
		if found {
			return id
		}
	}

	return 0
}

func getMappingFromLine(line string) (int, Mapping) {
	m := make(Mapping)
	parts := strings.Fields(line)
	id := util.ConvertStringToInt(strings.Trim(parts[1], ":"))
	parts = parts[2:]
	for i := 0; i <= len(parts)/2+1; i += 2 {
		word := strings.Trim(parts[i], ":,")
		if !util.IsNumber(word) {
			value := strings.Trim(parts[i+1], ":,")
			m[word] = util.ConvertStringToInt(value)
		}
	}

	return id, m
}
