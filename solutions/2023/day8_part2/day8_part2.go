package day8_part2

import (
	"strings"
)

func howManyStepsNodesEnd(lines []string) int {

	nodes := make(map[string][2]string)
	routes := make([]string, 0)
	for _, line := range lines[2:] {
		name := line[0:3]
		left := line[7:10]
		right := line[12:15]
		nodes[name] = [2]string{left, right}
		if strings.HasSuffix(name, "A") {
			routes = append(routes, name)
		}
	}

	//for _, dir := range lines[0] {
	//	for _, name := range routes {
	//		val, ok := nodes[name]
	//		if ok {
	//
	//		}
	//	}
	//}

	//numberOfSteps := 0
	//
	//countSteps := func(name string) int {
	//
	//}

	//outer:
	//for {
	//	for _, dir := range lines[0] {
	//		nextEntries := []string{}
	//		for _, name := range routes {
	//			val, ok := nodes[name]
	//			if ok {
	//				if dir == 'L' {
	//					v := val[0]
	//					nextEntries = append(nextEntries, v)
	//					//fmt.Printf("NODES: %v, %s = %s\n", nodes[name], val[0], val[1])
	//					//fmt.Printf("dir: %c, %q = %q\n", dir, name, v)
	//				} else {
	//					v := val[1]
	//					nextEntries = append(nextEntries, v)
	//					//fmt.Printf("NODES: %v, %s = %s\n", nodes[name], val[0], val[1])
	//					//fmt.Printf("dir: %c, %q = %q\n", dir, name, v)
	//				}
	//			}
	//		}
	//		routes = []string{}
	//		for _, v := range nextEntries {
	//			routes = append(routes, v)
	//		}
	//		numberOfSteps++
	//		fmt.Printf("Steps: %d\n", numberOfSteps)
	//		finished := 0
	//		for _, name := range nextEntries {
	//			if strings.HasSuffix(name, "Z") {
	//				finished++
	//			}
	//		}
	//		if finished == len(nextEntries) {
	//			return numberOfSteps
	//		}
	//		//if numberOfSteps >= 1000000 {
	//		//	break outer
	//		//}
	//		//time.Sleep(time.Second / 2)
	//	}
	//	//loops++
	//	//if loops >= 1 {
	//	//	break outer
	//	//}
	//}

	//return numberOfSteps
	return 0
}
