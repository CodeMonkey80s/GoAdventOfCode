package day6_part2

import "fmt"

func getAnswer(lines []string) int {

	ans := 0

	freq := make(map[rune]int)
	n := 0
	for i, line := range lines {
		if line != "" {
			for _, ch := range line {
				freq[ch]++
			}
			n++
		}

		if line == "" || i == len(lines)-1 {
			sum := 0
			for _, v := range freq {
				if v >= n {
					sum++
				}
			}
			ans += sum
			//fmt.Printf("freq: %+q, n: %d, sum: %d\n", freq, n, sum)
			freq = make(map[rune]int)
			n = 0
		}
	}

	fmt.Printf("Answer: %d\n", ans)
	return ans
}
