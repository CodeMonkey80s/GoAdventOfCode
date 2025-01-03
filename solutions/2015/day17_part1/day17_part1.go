package day17_part1

import (
	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	containers := make([]int, 0)
	for _, line := range lines {
		volume := util.ConvertStringToInt(line)
		containers = append(containers, volume)
	}
	perms := findCombos(containers, 150)

	return len(perms)
}

func findCombos(containers []int, target int) [][]int {
	combos := [][]int{}
	backtrack(containers, target, []int{}, &combos)
	return combos
}

func backtrack(containers []int, target int, currentCombo []int, combos *[][]int) {
	if target == 0 {
		// Found a valid combo
		*combos = append(*combos, append([]int{}, currentCombo...))
		return
	}
	if target < 0 || len(containers) == 0 {
		// Invalid combo
		return
	}
	// Explore all possible combinations
	for i := 0; i < len(containers); i++ {
		container := containers[i]
		remainingContainers := append([]int{}, containers[i+1:]...)
		backtrack(remainingContainers, target-container, append(currentCombo, container), combos)
	}
}
