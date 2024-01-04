package day19_part1

import (
	"fmt"
	"strings"
)

func getAnswer(lines []string) int {

	molecules := make(map[string]int)
	original := lines[len(lines)-1]

loop:
	for _, line := range lines[:len(lines)-2] {
		i := 0
		replacements := strings.Split(line, " => ")
		molecule := original
		for {
			m := molecule[i:]
			if len(m) == 0 {
				break
			}
			//fmt.Printf("i: %d, Searching in: %q => %q\n", i, molecule, m)
			idx := strings.Index(m, replacements[0])
			if idx == -1 {
				continue loop
			}
			t := i
			i += idx + len(replacements[0])
			s := strings.Replace(m, replacements[0], replacements[1], 1)
			s = molecule[:t] + s
			molecules[s]++
			//fmt.Printf("idx: %d, m: %q => %q\n", idx, m, s)
		}
		//fmt.Println("-----")
	}

	total := len(molecules)
	fmt.Printf("Output: %v\n", total)

	return total
}
