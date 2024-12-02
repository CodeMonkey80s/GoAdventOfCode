package day20_part1

func getAnswer(number int) int {

	input := number / 10
	houses := make([]int, input)
	houseNumber := input

	for i := 1; i < input; i++ {
		for j := i; j < input; j += i {
			houses[j] += i
			if houses[j] >= input && j < houseNumber {
				houseNumber = j
			}
		}

	}

	// fmt.Printf("house: %d\n", houseNumber)

	return houseNumber
}
