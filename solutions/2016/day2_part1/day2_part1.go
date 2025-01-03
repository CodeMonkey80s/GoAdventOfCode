package day2_part1

var board = [3][3]rune{
	{'1', '2', '3'},
	{'4', '5', '6'},
	{'7', '8', '9'},
}

func getAnswer(lines []string) string {
	sx := 1
	sy := 1
	code := ""
	for _, line := range lines {
		for i, ch := range line {
			end := i == len(line)-1
			switch ch {
			case 'U':
				if sy == 0 {
					continue
				}
				sy--
			case 'D':
				if sy == 2 {
					break
				}
				sy++
			case 'R':
				if sx == 2 {
					break
				}
				sx++
			case 'L':
				if sx == 0 {
					break
				}
				sx--
			}
			if end {
				code += string(board[sy][sx])
			}
		}
	}

	return code
}
