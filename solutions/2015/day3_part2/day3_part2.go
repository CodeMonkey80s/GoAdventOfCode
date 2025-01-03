package day3_part2

type House struct {
	X int
	Y int
}

func getAnswer(input string) int {
	houses := make(map[House]int)
	sx := 0
	sy := 0
	rx := 0
	ry := 0
	h := House{
		X: sx,
		Y: sy,
	}
	houses[h] += 2
	for i, char := range input {
		dx, dy := getDeltaCoords(char)
		if i%2 == 0 {
			sx = sx + dx
			sy = sy + dy
			h := House{
				X: sx,
				Y: sy,
			}
			houses[h]++
		} else {
			rx = rx + dx
			ry = ry + dy
			h := House{
				X: rx,
				Y: ry,
			}
			houses[h]++
		}
	}
	total := 0
	for _, v := range houses {
		if v >= 1 {
			total++
		}
	}
	return total
}

func getDeltaCoords(char rune) (int, int) {
	dx := 0
	dy := 0
	switch char {
	case '>':
		dx = 1
		dy = 0
	case '<':
		dx = -1
		dy = 0
	case '^':
		dx = 0
		dy = -1
	case 'v':
		dx = 0
		dy = 1
	}
	return dx, dy
}
