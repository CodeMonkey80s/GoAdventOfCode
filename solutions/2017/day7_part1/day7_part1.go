package day7_part1

import (
	"sort"
	"strings"

	"GoAdventOfCode/util"
)

type Program struct {
	Name     string
	Weight   int
	Children []Program
}

func getAnswer(lines []string) string {

	lines = sortInput(lines)
	programs := loadPrograms(lines)
	p := findRootProgram(programs)

	return p.Name
}

func sortInput(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		p1 := strings.Fields(lines[i])
		p2 := strings.Fields(lines[j])
		return p1[1] < p2[1]
	})
	return lines
}

func loadPrograms(lines []string) []Program {

	programs := make([]Program, 0)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		partLeft := strings.Fields(parts[0])

		partRight := []string{}
		if len(parts) > 1 {
			partRight = strings.Fields(parts[1])
		}
		name := partLeft[0]
		weight := util.ConvertStringToInt(strings.Trim(partLeft[1], "()"))

		children := make([]Program, 0)
		if len(partRight) > 1 {
			for _, childName := range partRight {
				name := strings.Trim(childName, ",")
				child := Program{
					Name: name,
				}
				children = append(children, child)
			}
		}

		p := Program{
			Name:     name,
			Weight:   weight,
			Children: children,
		}
		programs = append(programs, p)
	}

	return programs
}

// write code that will find root program
func findRootProgram(programs []Program) Program {
	childMap := make(map[string]bool)

	// Create a map of all child programs
	for _, program := range programs {
		for _, child := range program.Children {
			childMap[child.Name] = true
		}
	}

	// Find the program that is not listed as a child in any other program
	for _, program := range programs {
		if _, ok := childMap[program.Name]; !ok {
			return program
		}
	}

	return Program{} // Return an empty program if root program is not found
}
