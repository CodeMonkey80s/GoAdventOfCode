package day2_part2

import (
	"slices"
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

	for n, line := range lines {

		parts := strings.Fields(line)
		numbers[n] = make([]int, 0, len(parts))

		for i := 0; i < len(parts); i++ {
			num, _ := strconv.Atoi(parts[i])
			numbers[n] = append(numbers[n], num)
		}
	}

	checkIsSafe := func(nums []int) bool {
		a := 0
		b := 1
		order := orderUnknown
		for i := 0; i < len(nums)-1; i++ {
			switch {
			case nums[a] == nums[b]:
				return false
			case nums[b] > nums[a] && nums[b]-nums[a] >= 1 && nums[b]-nums[a] <= 3:
				if order == orderDes {
					return false
				}
				order = orderAsc
			case nums[a] > nums[b] && nums[a]-nums[b] >= 1 && nums[a]-nums[b] <= 3:
				if order == orderAsc {
					return false
				}
				order = orderDes
			default:
				return false
			}
			a++
			b++
		}

		if order != orderUnknown {
			return true
		}

		return false
	}

	problemDampener := func(numsOrg []int) bool {
		for i := 0; i < len(numsOrg); i++ {
			nums := make([]int, len(numsOrg))
			copy(nums, numsOrg)
			nums = slices.Delete(nums, i, i+1)
			// fmt.Printf("i: %d\n", i)
			// fmt.Printf("org: %+v\n", numsOrg)
			// fmt.Printf("num: %+v\n", nums)
			isSafe := checkIsSafe(nums)
			// fmt.Printf("isSafe: %t\n", isSafe)
			if isSafe {
				return true
			}
		}

		return false
	}

	safe := 0
	for _, nums := range numbers {
		isSafe := problemDampener(nums)
		// fmt.Printf("nums: %+v, %t\n", nums, isSafe)
		if isSafe {
			safe++
		}
	}

	return safe
}
