package day5_part1

func getAnswer(lines []string) int {

	ans := 0
	for _, line := range lines {
		val := findRow(line)*8 + findColumn(line)
		ans = max(ans, val)
	}
	return ans
}

func findRow(boardingPass string) int {
	lower := 0
	upper := 127

	for i := 0; i < 7; i++ {
		mid := (lower + upper) / 2
		if boardingPass[i] == 'F' {
			upper = mid
		} else if boardingPass[i] == 'B' {
			lower = mid + 1
		}
	}

	return lower
}

func findColumn(boardingPass string) int {
	lower := 0
	upper := 7

	for i := 7; i < 10; i++ {
		mid := (lower + upper) / 2
		if boardingPass[i] == 'L' {
			upper = mid
		} else if boardingPass[i] == 'R' {
			lower = mid + 1
		}
	}

	return lower
}
