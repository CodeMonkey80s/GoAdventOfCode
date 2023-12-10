package day8_part1

import (
	"log"
)

type MapNode struct {
	Left  string
	Right string
}

const (
	nodeStartID = "AAA"
	nodeEndID   = "ZZZ"
)

func howManySteps(lines []string) int {

	steps := lines[0]
	nodes := make(map[string]MapNode)
	for _, line := range lines[2:] {
		id := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[id] = MapNode{
			Left:  left,
			Right: right,
		}
	}

	numberOfSteps := 0
	id := nodeStartID
	next := ""
outer:
	for {
		for _, dir := range steps {
			switch dir {
			case 'L':
				v, ok := nodes[id]
				if ok {
					next = v.Left
				} else {
					log.Panicf("node %q not found", id)
				}
			case 'R':
				v, ok := nodes[id]
				if ok {
					next = v.Right
				} else {
					log.Panicf("node %q not found", id)
				}
			}
			numberOfSteps++
			if next == nodeEndID {
				break outer
			}
			id = next
		}
	}

	return numberOfSteps
}
