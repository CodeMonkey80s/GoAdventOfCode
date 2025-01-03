package day25_part1

type schematic struct {
	t       string
	heights [5]int
}

func getAnswer(input []string) int {
	overlapping := 0
	keys, locks := loadSchematics(input)
	for _, k := range keys {
		for _, l := range locks {
			ok := 0
			for i := 0; i < 5; i++ {
				if k.heights[i]+l.heights[i] < 6 {
					ok++
				}
			}
			if ok == 5 {
				overlapping++
			}
		}
	}

	return overlapping
}

func loadSchematics(input []string) ([]schematic, []schematic) {
	var keys []schematic
	var locks []schematic
	for i := 0; i < len(input)-7; i += 7 {
		a := 1
		b := 7
		s := schematic{
			t:       "lock",
			heights: [5]int{},
		}
		if input[i] == "....." {
			s.t = "key"
			a = 0
			b = 6
		}

		for x := 0; x < 5; x++ {
			count := 0
			for y := a; y < b; y++ {
				ch := input[i+y][x]
				if ch == '#' {
					count++
				}
			}
			s.heights[x] += count
		}

		switch {
		case s.t == "key":
			keys = append(keys, s)
		case s.t == "lock":
			locks = append(locks, s)
		}

		i++
	}

	return keys, locks
}
