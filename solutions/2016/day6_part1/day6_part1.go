package day6_part1

func mostCommonCharacters(lines []string) string {
	ans := ""
	for i := 0; i < len(lines[0]); i++ {
		freq := make(map[rune]int)
		for _, line := range lines {
			char := rune(line[i])
			freq[char]++
		}
		m := 0
		key := rune(0)
		for k, v := range freq {
			if v > m {
				m = v
				key = k
			}
		}
		clear(freq)
		ans += string(key)
	}

	return ans
}
