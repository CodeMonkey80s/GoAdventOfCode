package day13_part2

import (
	"maps"
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

type person struct {
	name  string
	rules map[string]int
}

func getAnswer(lines []string) int {

	people := make(map[string]person)

	happiness := 0
	for _, line := range lines {

		parts := strings.Fields(line)

		name := parts[0]
		happinessChange := util.ConvertStringToInt(parts[3])
		if parts[2] == "lose" {
			happinessChange = -happinessChange
		}

		_, ok := people[name]
		if !ok {
			p := person{
				name:  name,
				rules: make(map[string]int),
			}
			people[name] = p
			people[name].rules["Me"] = 0

			me := person{
				name:  "Me",
				rules: make(map[string]int),
			}
			people["Me"] = me
		}

		guest := strings.TrimRight(parts[10], ".")
		people[name].rules[guest] = happinessChange

	}

	names := slices.Collect(maps.Keys(people))
	for _, name := range names {
		if name != "Me" {
			people["Me"].rules[name] = 0
		}
	}

	// util.PrintOrderedMap("people", people)

	util.PermString(names, func(names []string) {

		// util.PrintSlice("names", names)

		diff := 0
		for i := 0; i < len(names); i++ {

			name := names[i]

			a := i - 1
			if a < 0 {
				a = len(names) - 1
			}
			guestL := names[a]
			diffL := people[name].rules[guestL]

			b := i + 1
			if b > len(names)-1 {
				b = 0
			}
			guestR := names[b]
			diffR := people[name].rules[guestR]

			diff += diffL + diffR

		}

		// fmt.Printf("diff: %d\n", diff)
		happiness = max(happiness, diff)
		// fmt.Printf("happiness: %d\n", happiness)
		// fmt.Println()
	})

	return happiness
}
