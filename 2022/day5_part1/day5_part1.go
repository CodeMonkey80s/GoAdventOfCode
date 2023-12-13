package day5_part1

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

type Stack struct {
	Items []string
}

func (st *Stack) Initialize() {
	st.Items = make([]string, 0)
}

func (st *Stack) Push(item string) {
	st.Items = append(st.Items, item)
}

func (st *Stack) Pop() string {
	m := len(st.Items) - 1
	item := st.Items[m]
	st.Items = st.Items[:m]
	return item
}

func (st *Stack) Last() string {
	m := len(st.Items) - 1
	return st.Items[m]
}

type Stacks struct {
	Index map[int]*Stack
}

func (st *Stacks) Initialize(n int) {
	st.Index = make(map[int]*Stack)
	for i := 1; i <= n; i++ {
		stack := Stack{}
		stack.Initialize()
		st.Index[i] = &stack
	}
}

var CrateYard Stacks

func finalCratesPositions(lines []string) string {
	idx := findBreakingIndex(lines)
	setupCrates(lines[:idx])
	moveCrates(lines[idx+1:])
	return getCratesPlacement()
}

func findBreakingIndex(lines []string) int {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	return -1
}

func setupCrates(lines []string) {
	stacksLine := lines[len(lines)-1]
	stacksLine = strings.Replace(stacksLine, " ", "", -1)
	stacksLine = strings.TrimRight(stacksLine, " ")
	stacks := util.ConvertByteToInt(stacksLine[len(stacksLine)-1])
	CrateYard.Initialize(stacks)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i][0:]
		for j := 0; j < stacks; j++ {
			crate := line[j*4 : (j*4)+3]
			if crate == "   " {
				continue
			}
			CrateYard.Index[j+1].Push(crate)
		}
	}
}

func moveCrates(lines []string) {
	for _, line := range lines {
		parts := strings.Fields(line)
		n := util.ConvertStringToInt(parts[1])
		a := util.ConvertStringToInt(parts[3])
		b := util.ConvertStringToInt(parts[5])
		for i := 0; i < n; i++ {
			crate := CrateYard.Index[a].Pop()
			CrateYard.Index[b].Push(crate)
		}
	}
}

func getCratesPlacement() string {
	//printStacking()
	crates := ""
	for i := 1; i <= len(CrateYard.Index); i++ {
		crates += string(CrateYard.Index[i].Last()[1])
	}
	return crates
}

func printStacking() {
	for i := 1; i <= len(CrateYard.Index); i++ {
		stack, ok := CrateYard.Index[i]
		if ok {
			fmt.Printf("%d -> ", i)
			for _, item := range stack.Items {
				fmt.Printf("%v", item)
			}
			fmt.Println()
		}
	}
}
