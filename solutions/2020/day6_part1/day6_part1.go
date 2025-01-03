package day6_part1

import "fmt"

func getAnswer(lines []string) int {

	ans := 0
	freq := make(map[rune]int)
	for i, line := range lines {
		for _, ch := range line {
			freq[ch]++
		}
		if line == "" || i == len(lines)-1 {
			ans += len(freq)
			freq = make(map[rune]int)
		}
	}

	fmt.Printf("Answer: %d\n", ans)
	return ans
}
