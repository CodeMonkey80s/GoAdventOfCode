package day3_part1

type House struct {
	X int
	Y int
}

func getAnswer(input string) int {
	houses := make(map[House]int)
	x := 0
	y := 0
	h := House{
		X: x,
		Y: y,
	}
	houses[h]++
	for _, char := range input {
		switch char {
		case '>':
			x++
		case '<':
			x--
		case '^':
			y--
		case 'v':
			y++
		}
		h := House{
			X: x,
			Y: y,
		}
		houses[h]++
	}
	total := 0
	for _, v := range houses {
		if v >= 1 {
			total++
		}
	}
	//fmt.Printf("houses: %+v\n", houses)
	return total
}
