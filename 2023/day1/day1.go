package day1

import (
	"strconv"
	"strings"
)

func sumOfAllTheCalibrationValues_part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		nums := ""
		for _, ch := range line {
			if ch >= '0' && ch <= '9' {
				nums += string(ch)
			}
		}
		if len(nums) == 1 {
			nums = nums + nums
		}
		if len(nums) >= 3 {
			nums = string(nums[0]) + string(nums[len(nums)-1])
		}
		val, err := strconv.Atoi(nums)
		if err != nil {
			panic(err)
		}
		sum += val
	}

	return sum
}

func sumOfAllTheCalibrationValues_part2(lines []string) int {
	sum := 0
	nums := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for _, line := range lines {

		num1 := ""
		num2 := ""

		for i := 0; i < len(line); i++ {
			ch := line[i]
			if num1 == "" && ch >= '0' && ch <= '9' {
				num1 = string(ch)
				break
			}
			if num1 == "" {
				for j := 0; j < len(nums); j++ {
					if strings.HasPrefix(line[i:], nums[j]) {
						num1 = string('0' + byte(j) + 1)
						break
					}
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			ch := line[i]
			if num2 == "" && ch >= '0' && ch <= '9' {
				num2 = string(ch)
				break
			}
			if num2 == "" {
				for j := 0; j < len(nums); j++ {
					if strings.HasSuffix(line[:i+1], nums[j]) {
						num2 = string('0' + byte(j) + 1)
						break
					}
				}
			}
		}

		if num1 != "" && num2 != "" {
			number := num1 + num2
			val, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			sum += val
		}
	}

	return sum
}
