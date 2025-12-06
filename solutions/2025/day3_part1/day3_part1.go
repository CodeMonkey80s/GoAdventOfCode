package day3_part1

func getAnswerForPart1(lines []string) int {

	var joltage int

	for _, s := range lines {

		var d1 int
		var idx int
		for i := 0; i < len(s)-1; i++ {
			v := int(s[i] - '0')
			if v > d1 {
				d1 = v
				idx = i
			}
		}

		var d2 int
		for i := idx + 1; i < len(s); i++ {
			v := int(s[i] - '0')
			d2 = max(d2, v)
		}

		joltage += d1*10 + d2
	}

	return joltage
}
