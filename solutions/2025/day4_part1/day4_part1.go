package day4_part1

func getAnswerForPart1(lines []string) int {

	directions := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	isCountAsRoll := func(y int, x int, lines []string) bool {
		if lines[y][x] != '@' {
			return false
		}

		var counter int
		for i := 0; i < len(directions); i++ {
			dy := directions[i][0]
			dx := directions[i][1]
			tx := x + dx
			ty := y + dy
			if tx < 0 || tx >= len(lines[0]) {
				continue
			}
			if ty < 0 || ty >= len(lines) {
				continue
			}
			if lines[ty][tx] == '@' {
				counter++
			}
		}

		if counter < 4 {
			return true
		}

		return false
	}

	var rolls int

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			ok := isCountAsRoll(y, x, lines)
			if ok {
				rolls++
			}
		}
	}

	return rolls
}
