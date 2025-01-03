package day4_part1

const (
	charIgnore = ' '
)

var masks = [][]string{
	{
		"XMAS",
	},
	{
		"SAMX",
	},
	{
		"X",
		"M",
		"A",
		"S",
	},
	{
		"S",
		"A",
		"M",
		"X",
	},
	{
		"X   ",
		" M  ",
		"  A ",
		"   S",
	},
	{
		"   X",
		"  M ",
		" A  ",
		"S   ",
	},
	{
		"S   ",
		" A  ",
		"  M ",
		"   X",
	},
	{
		"   S",
		"  A ",
		" M  ",
		"X   ",
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
