package day5_part1

import (
	"slices"
	"strconv"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {

	numbersIdx := 0
	ordering := make([][2]int, 0)
	for i, line := range lines {
		if line == "" {
			numbersIdx = i + 1
			continue
		}
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			v1, _ := strconv.Atoi(parts[0])
			v2, _ := strconv.Atoi(parts[1])
			order := [2]int{v1, v2}
			ordering = append(ordering, order)
		}
	}

	checkNumbers := func(orders [][2]int, numbers []int) (int, bool) {
		correct := make([]bool, len(numbers))
		total := 0
		count := 0
		// fmt.Printf("numbers: %+v\n", numbers)
		for idx := range numbers {
		inner:
			for _, order := range orders {
				if !slices.Contains(numbers, order[0]) || !slices.Contains(numbers, order[1]) {
					continue
				}

				total++

				idx1 := 0
				idx2 := len(numbers) - 1
				for i := 0; i < len(numbers); i++ {
					if numbers[i] == order[0] {
						idx1 = i
						break
					}
				}
				for i := len(numbers) - 1; i >= 0; i-- {
					if numbers[i] == order[1] {
						idx2 = i
						break
					}
				}
				if idx1 < idx2 {
					// fmt.Printf("order: %+v\n", order)
					correct[idx] = true
					count++
					continue inner
				}
			}
		}

		// fmt.Printf("correct: %+v\n", correct)
		if count == total {
			middle := numbers[len(numbers)/2]
			// fmt.Printf("middle: %d\n", middle)
			return middle, true
		}

		return 0, false
	}

	getNumbers := func(line string) []int {
		var numbers []int
		parts := strings.Split(line, ",")
		for _, part := range parts {
			num := util.ConvertStringToInt(part)
			numbers = append(numbers, num)
		}

		return numbers
	}

	sum := 0
	for _, line := range lines[numbersIdx:] {
		numbers := getNumbers(line)
		val, ok := checkNumbers(ordering, numbers)
		if ok {
			// fmt.Printf("line: %q correct order\n", line)
			sum += val
		}
	}

	// fmt.Printf("idx: %d\n", numbersIdx)
	// fmt.Printf("ordering: %+v\n", ordering)
	// fmt.Printf("sum: %d\n", sum)

	return sum
}
