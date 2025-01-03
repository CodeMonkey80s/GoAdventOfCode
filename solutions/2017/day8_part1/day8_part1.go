package day8_part1

import (
	"fmt"
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

type Registers map[string]int

func getAnswer(lines []string) int {

	registers := make(Registers)

	for _, line := range lines {

		// 0    1         2    3  4    5     6
		// [r1] [inc|dec] [v1] if [r2] [cmp] [v2]

		parts := strings.Fields(line)
		r1 := parts[0]
		r2 := parts[4]
		cmd := parts[1]
		cmp := parts[5]
		v1 := util.ConvertStringToInt(parts[2])
		v2 := util.ConvertStringToInt(parts[6])

		if _, ok := registers[r1]; !ok {
			registers[r1] = 0
		}
		if _, ok := registers[r2]; !ok {
			registers[r2] = 0
		}

		switch {

		case cmd == "inc" && cmp == "==":
			if registers[r2] == v2 {
				registers[r1] += v1
			}
		case cmd == "inc" && cmp == "!=":
			if registers[r2] != v2 {
				registers[r1] += v1
			}
		case cmd == "inc" && cmp == ">":
			if registers[r2] > v2 {
				registers[r1] += v1
			}
		case cmd == "inc" && cmp == ">=":
			if registers[r2] >= v2 {
				registers[r1] += v1
			}
		case cmd == "inc" && cmp == "<":
			if registers[r2] < v2 {
				registers[r1] += v1
			}
		case cmd == "inc" && cmp == "<=":
			if registers[r2] <= v2 {
				registers[r1] += v1
			}

		case cmd == "dec" && cmp == "==":
			if registers[r2] == v2 {
				registers[r1] -= v1
			}
		case cmd == "dec" && cmp == "!=":
			if registers[r2] != v2 {
				registers[r1] -= v1
			}
		case cmd == "dec" && cmp == "<":
			if registers[r2] < v2 {
				registers[r1] -= v1
			}
		case cmd == "dec" && cmp == "<=":
			if registers[r2] <= v2 {
				registers[r1] -= v1
			}
		case cmd == "dec" && cmp == ">":
			if registers[r2] > v2 {
				registers[r1] -= v1
			}
		case cmd == "dec" && cmp == ">=":
			if registers[r2] >= v2 {
				registers[r1] -= v1
			}
		}
	}

	maxValue := 0
	//printRegisters(registers)
	for _, v := range registers {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}

func printRegisters(registers Registers) {
	keys := make([]string, 0, len(registers))
	for k := range registers {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, v := range keys {
		fmt.Printf("%v => \"%+v\"\n", v, registers[v])
	}
}
