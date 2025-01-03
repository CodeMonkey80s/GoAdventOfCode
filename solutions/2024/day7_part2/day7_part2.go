package day7_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	sum := 0
outer:
	for _, line := range lines {
		parts := strings.Fields(line)

		total := util.ConvertStringToInt(strings.TrimSuffix(parts[0], ":"))

		var numbers []int
		for _, num := range parts[1:] {
			numbers = append(numbers, util.ConvertStringToInt(num))
		}

		size := len(numbers) - 1
		combinations := getCombinations([]rune{'+', '*', '|'}, size)
		for _, combo := range combinations {
			val := checkCombination(numbers, total, combo)
			if val == total {
				sum += total
				continue outer
			}
		}
	}

	return sum
}

func getCombinations(letters []rune, length int) []string {
	if length == 0 {
		return []string{""}
	}

	var result []string
	subCombinations := getCombinations(letters, length-1)

	for _, letter := range letters {
		for _, sub := range subCombinations {
			result = append(result, string(letter)+sub)
		}
	}

	return result
}

func checkCombination(numbers []int, total int, combo string) int {
	sum := numbers[0]
	for i, num := range numbers[1:] {
		switch combo[i] {
		case '+':
			sum += num
		case '*':
			sum *= num
		case '|':
			s1 := util.ConvertIntToString(sum)
			s2 := util.ConvertIntToString(num)
			s := s1 + s2
			n := util.ConvertStringToInt(s)
			sum = n
		}
	}

	return sum
}
