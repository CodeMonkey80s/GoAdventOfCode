package day24_part1

type point struct {
	q int
	r int
	s int
}

func getAnswer(input []string) int {

	grid := make(map[point]rune)

	flip := func(p point) {
		color, ok := grid[p]
		if !ok {
			grid[p] = 'B'
			return
		}

		if color == 'W' {
			grid[p] = 'B'
			return
		}

		grid[p] = 'W'
	}

	var pointer point

loop:
	for _, line := range input {

		pointer.s = 0
		pointer.q = 0
		pointer.r = 0

		i := 0
		for {
			switch {
			case line[i] == 'n' && line[i+1] == 'w':
				pointer.r--
				pointer.s++
				i += 2
			case line[i] == 'n' && line[i+1] == 'e':
				pointer.r--
				pointer.q++
				i += 2
			case line[i] == 's' && line[i+1] == 'e':
				pointer.r++
				pointer.s--
				i += 2
			case line[i] == 's' && line[i+1] == 'w':
				pointer.r++
				pointer.q--
				i += 2
			case line[i] == 'e':
				pointer.q++
				pointer.s--
				i++
			case line[i] == 'w':
				pointer.q--
				pointer.s++
				i++
			default:
				panic("something is wrong!")
			}

			if i >= len(line) {
				flip(pointer)
				continue loop
			}
		}
	}

	tiles := 0
	for _, color := range grid {
		if color == 'B' {
			tiles++
		}
	}

	return tiles
}
