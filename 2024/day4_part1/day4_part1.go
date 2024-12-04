package day4_part1

var masksHorizontal = [][]string{
	{
		"XMAS",
	},
	{
		"SAMX",
	},
}

var masksVertical = [][]string{
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
}

var masksDiagonal = [][]string{
	{
		"X...",
		".M..",
		"..A.",
		"...S",
	},
	{
		"...X",
		"..M.",
		".A..",
		"S...",
	},
	{
		"S...",
		".A..",
		"..M.",
		"...X",
	},
	{
		"...S",
		"..A.",
		".M..",
		"X...",
	},
}

func getAnswer(lines []string) int {

	sum := 0

	sum += findMask(masksHorizontal[0], lines)
	sum += findMask(masksHorizontal[1], lines)

	sum += findMask(masksVertical[0], lines)
	sum += findMask(masksVertical[1], lines)

	sum += findMask(masksDiagonal[0], lines)
	sum += findMask(masksDiagonal[1], lines)
	sum += findMask(masksDiagonal[2], lines)
	sum += findMask(masksDiagonal[3], lines)

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
					if mask[dy][dx] != '.' {
						total++
					}

					if mask[dy][dx] == lines[y+dy][x+dx] {
						found++
					}

				}
			}

			// fmt.Printf("total: %d, found: %d\n", total, found)

			if found == total {
				count++
			}

		}
	}

	return count
}
