package day6_part2

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func howManyCycles(input string) int {

	blocks := initialize(input)
	cycles := process(blocks)

	fmt.Printf("Cyces: %d\n", cycles)

	return cycles
}

func initialize(input string) []int {
	blocks := make([]int, 0)
	parts := strings.Fields(input)
	for _, s := range parts {
		blocks = append(blocks, util.ConvertStringToInt(s))
	}

	return blocks
}

func process(blocks []int) int {
	idx := 0
	cycles := 0
	next := 0
	configurations := make(map[string]bool)
	for {
		idx = findHighBlock(blocks)
		blocks = redistribute(blocks, idx)
		h := hash(blocks)
		cycles++
		if _, ok := configurations[h]; ok {
			clear(configurations)
			if next == 1 {
				break
			}
			next++
			cycles = 0
		}
		configurations[h] = true
	}
	return cycles
}

func redistribute(blocks []int, idx int) []int {
	n := 0
	i := idx + 1
	if i > len(blocks)-1 {
		i = 0
	}
	c := blocks[idx]
	blocks[idx] = 0
	for {
		blocks[i]++
		i++
		n++
		if n >= c {
			break
		}
		if i > len(blocks)-1 {
			i = 0
		}
	}
	return blocks
}

func hash(blocks []int) string {
	sb := strings.Builder{}
	for _, v := range blocks {
		sb.WriteString(util.ConvertIntToString(v))
	}
	h := sb.String()
	return h
}

func findHighBlock(blocks []int) int {
	idx := 0
	maxVal := 0
	for i, v := range blocks {
		if v > maxVal {
			maxVal = v
			idx = i
		}
	}
	return idx
}
