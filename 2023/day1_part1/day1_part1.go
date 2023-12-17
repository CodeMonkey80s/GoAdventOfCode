package day1_part1

import (
	"GoAdventOfCode/util"
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
		val := util.ConvertStringToInt(nums)
		sum += val
	}

	return sum
}
