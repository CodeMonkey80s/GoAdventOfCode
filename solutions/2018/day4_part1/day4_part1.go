package day4_part1

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"GoAdventOfCode/util"
)

type Guard struct {
	TimeSleep int
	Frequency map[int]int
}

func getAnswer(lines []string) int {

	sortedLines := sortInput(lines)

	guards := make(map[string]Guard)

	id := ""
	t1 := ""
	t2 := ""
	for i, line := range sortedLines {
		t := line[0:18]
		parts := strings.Fields(line[18:])
		switch {
		case parts[0] == "Guard":
			id = parts[1]
		case parts[0] == "falls":
			t1 = t
		case parts[0] == "wakes":
			t2 = t
		}

		if i == len(sortedLines)-1 || id != "" && t1 != "" && t2 != "" {
			guards = calculateSleepTimeBetween(guards, id, t1, t2)

			t1 = ""
			t2 = ""
		}
	}

	//printGuards(guards)

	id = ""
	minutes := 0
	for k, g := range guards {
		if g.TimeSleep > minutes {
			id = k
			minutes = g.TimeSleep
		}
	}

	minute := 0
	value := 0
	for k, v := range guards[id].Frequency {
		if v > value {
			value = v
			minute = k
		}
	}

	ans := minute * util.ConvertStringToInt(id[1:])
	fmt.Printf("Most Time Asleep ID: %v (%d Minutes)\n", id, minutes)
	fmt.Printf("Minute: %v\n", minute)
	fmt.Printf("Output: %v\n", ans)

	return ans
}

func calculateSleepTimeBetween(guards map[string]Guard, id string, t1 string, t2 string) map[string]Guard {
	date := t1[1:11]
	g := Guard{
		TimeSleep: 0,
		Frequency: make(map[int]int),
	}
	currentGuard, ok := guards[id]
	if ok {
		g = currentGuard
	}
	for n := 0; n <= 59; n++ {
		m := fmt.Sprintf("00:%02d", n)
		d := fmt.Sprintf("[%s %s]", date, m)
		if d >= t1 && d < t2 {
			g.Frequency[n]++
			g.TimeSleep++
		}
	}
	guards[id] = g
	return guards
}

func sortInput(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	return lines
}

func printGuards(guards map[string]Guard) {
	keys := make([]string, 0, len(guards))
	for k, _ := range guards {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, v := range keys {
		fmt.Printf("%v %v\n", v, guards[v])
	}
}
