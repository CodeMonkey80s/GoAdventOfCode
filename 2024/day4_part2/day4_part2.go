package day4_part2

const (
	charIgnore = ' '
)

var masks = [][]string{
	{
		"M S",
		" A ",
		"M S",
	},
	{
		"S S",
		" A ",
		"M M",
	},
	{
		"S M",
		" A ",
		"S M",
	},
	{
		"M M",
		" A ",
		"S S",
	},
}

func getAnswer(lines []string) int {

	sum := 0
	for _, mask := range masks {
		sum += findMask(mask, lines)
	}

	return sum
}

func findMask(mask []string, lines []string) int {

	count := 0
	for y := 0; y <= len(lines)-len(mask); y++ {
		for x := 0; x <= len(lines[0])-len(mask[0]); x++ {

			found := 0
			total := 0
			for dy := 0; dy < len(mask); dy++ {
				for dx := 0; dx < len(mask[0]); dx++ {
					if mask[dy][dx] != charIgnore {
						total++
					}

					if mask[dy][dx] == lines[y+dy][x+dx] {
						found++
					}

				}
			}

			if found == total {
				count++
			}

		}
	}

	return count
}
