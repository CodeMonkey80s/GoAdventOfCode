package day9_part2

import (
	"fmt"
	"strings"

	"GoAdventOfCode/util"
)

func getLongestRoute(lines []string) int {
	citiesFrequency := make(map[string]int)
	listOfCities := make([]string, 0)
	citiesToDistance := make(map[string]int)
	for _, line := range lines {
		cities, distance := getCities(line)
		id1 := getID(cities)
		citiesToDistance[id1] = distance
		citiesFrequency[cities[0]] = 0
		citiesFrequency[cities[1]] = 0
		id2 := getID([]string{cities[1], cities[0]})
		citiesToDistance[id2] = distance

	}

	for city := range citiesFrequency {
		listOfCities = append(listOfCities, city)
	}

	longestDistance := 0
	citiesPermutations := getPermutations(listOfCities)

	for _, cities := range citiesPermutations {
		distance := 0
		a := 0
		b := 1
		for n := 0; n < len(cities)-1; n++ {
			route := getID([]string{cities[a], cities[b]})
			v, ok := citiesToDistance[route]
			if ok {
				distance += v
				a++
				b++
			}
		}
		longestDistance = max(longestDistance, distance)
	}

	fmt.Printf("Output: %d\n", longestDistance)

	return longestDistance
}

func getCities(line string) ([]string, int) {
	parts := strings.Fields(line)
	distance := util.ConvertStringToInt(parts[4])
	cities := []string{parts[0], parts[2]}
	return cities, distance
}

func getID(cities []string) string {
	return strings.Join(cities, "->")
}

func getPermutations(elements []string) [][]string {
	permutations := [][]string{}
	if len(elements) == 1 {
		permutations = [][]string{elements}
		return permutations
	}
	for i := range elements {
		el := make([]string, len(elements))
		copy(el, elements)
		for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
			permutations = append(permutations, append([]string{elements[i]}, perm...))
		}
	}
	return permutations
}
