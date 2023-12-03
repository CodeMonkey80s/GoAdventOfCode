package day1part1

import (
	"strconv"
)

func sumOfAllTheCalibrationValues(lines []string) int {
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
