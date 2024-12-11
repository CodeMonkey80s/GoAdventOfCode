package day11_part2

import (
	"GoAdventOfCode/util"
)

func getAnswer(stones []int, numberOfBlinks int) int {

	cache := make(map[int]int)
	for _, stone := range stones {
		cache[stone]++
	}

	for i := 0; i < numberOfBlinks; i++ {
		cache = blink(cache)
	}

	count := 0
	for _, n := range cache {
		count += n
	}
	return count
}

func blink(cacheOld map[int]int) map[int]int {

	cacheNew := make(map[int]int)
	for stone, count := range cacheOld {

		if stone == 0 {
			cacheNew[1] += count
			continue
		}

		digits, ok := hasEvenDigits(stone)
		if ok {
			num1 := digits[0 : len(digits)/2]
			num2 := digits[len(digits)/2:]
			v1 := util.ConvertStringToInt(num1)
			v2 := util.ConvertStringToInt(num2)
			cacheNew[v1] += count
			cacheNew[v2] += count
			continue
		}

		cacheNew[stone*2024] += count
	}

	return cacheNew
}

func hasEvenDigits(stone int) (string, bool) {
	digits := util.ConvertIntToString(stone)
	if len(digits)%2 == 0 {
		return digits, true
	}

	return "", false
}
