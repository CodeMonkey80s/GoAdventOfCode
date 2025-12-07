package day7_part1

func getAnswerForPart1(lines []string) int {
	var totalSplits int

	beams := make(map[int]struct{})

	for y, line := range lines {
		if y%2 != 0 {
			continue
		}
		var newBeams []int
		for x := 0; x < len(lines[0]); x++ {
			ch := line[x]
			_, ok := beams[x]
			switch {
			case ch == '^' && ok:
				delete(beams, x)
				newBeams = append(newBeams, x-1)
				newBeams = append(newBeams, x+1)
				totalSplits++
			case ch == 'S':
				newBeams = append(newBeams, x)
			}
		}

		for _, x := range newBeams {
			beams[x] = struct{}{}
		}
	}

	return totalSplits
}
