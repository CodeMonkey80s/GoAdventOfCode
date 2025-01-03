package day2_part1

import (
	"strconv"
	"strings"
)

const (
	orderUnknown = 0
	orderAsc     = 1
	orderDes     = 2
)

func getAnswer(lines []string) int {

	numbers := make([][]int, len(lines))

	safe := 0
	for n, line := range lines {

		parts := strings.Fields(line)
		numbers[n] = make([]int, 0, len(parts))

		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			numbers[n] = append(numbers[n], num)
		}
	}

outer:
	for _, nums := range numbers {
		a := 0
		b := 1
		order := orderUnknown
		for i := 0; i < len(nums)-1; i++ {
			switch {
			case nums[a] == nums[b]:
				continue outer
			case nums[b] > nums[a] && nums[b]-nums[a] >= 1 && nums[b]-nums[a] <= 3:
				if order == orderDes {
					continue outer
				}
				order = orderAsc
			case nums[a] > nums[b] && nums[a]-nums[b] >= 1 && nums[a]-nums[b] <= 3:
				if order == orderAsc {
					continue outer
				}
				order = orderDes
			default:
				continue outer
			}
			a++
			b++
		}

		if order != orderUnknown {
			safe++
		}
	}

	return safe
}
