package day20_part2

func getAnswer(number int) int {

	input := number / 10
	houses := make([]int, input)
	houseNumber := input

	for i := 1; i < input; i++ {
		visits := 0
		for j := i; j < input; j += i {
			houses[j] = houses[j] + i*11
			if houses[j] >= input*10 && j < houseNumber {
				houseNumber = j
			}
			visits++
			if visits == 50 {
				break
			}
		}

	}

	// fmt.Printf("house: %d\n", houseNumber)

	return houseNumber
}
