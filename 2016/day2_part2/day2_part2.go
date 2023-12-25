package day2_part2

const (
	size = 5
)

var board = [size][size]rune{
	{'.', '.', '1', '.', '.'},
	{'.', '2', '3', '4', '.'},
	{'5', '6', '7', '8', '9'},
	{'.', 'A', 'B', 'C', '.'},
	{'.', '.', 'D', '.', '.'},
}

func getAnswer(lines []string) string {
	sx := 0
	sy := 2
	code := ""
	for _, line := range lines {
		for i, ch := range line {
			end := i == len(line)-1
			switch ch {
			case 'U':
				if ok := canMove(sx, sy-1); !ok {
					break
				}
				sy--
			case 'D':
				if ok := canMove(sx, sy+1); !ok {
					break
				}
				sy++
			case 'R':
				if ok := canMove(sx+1, sy); !ok {
					break
				}
				sx++
			case 'L':
				if ok := canMove(sx-1, sy); !ok {
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

func canMove(sx int, sy int) bool {
	switch {
	case sx < 0:
		return false
	case sx > size-1:
		return false
	case sy > size-1:
		return false
	case sy < 0:
		return false
	}

	return board[sy][sx] != '.'
}
