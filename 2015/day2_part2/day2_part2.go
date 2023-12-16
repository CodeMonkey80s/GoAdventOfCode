package day2_part2

import (
	"slices"
	"strings"

	"GoAdventOfCode/util"
)

func getTotal(lines []string) int {
	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "x")
		a := util.ConvertStringToInt(parts[0])
		b := util.ConvertStringToInt(parts[1])
		c := util.ConvertStringToInt(parts[2])
		total += getRibbon([]int{a, b, c})
	}
	return total
}

func getRibbon(nums []int) int {
	slices.Sort(nums)
	m := nums[0] + nums[0] + nums[1] + nums[1]
	m += nums[0] * nums[1] * nums[2]
	return m
}
