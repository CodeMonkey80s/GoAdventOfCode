package day8_part1

import (
	"fmt"
	"math"
	"strings"
)

func getAnswer(input string, w int, h int) int {

	n := w * h

	layers := make([]string, 0)
	for i, v := range input {
		if i%n == 0 {
			layers = append(layers, "")
		}
		layers[len(layers)-1] += string(v)
	}

	//printLayers(layers)

	id := 0
	value := math.MaxUint32
	for number, layer := range layers {
		countZeros := strings.Count(layer, "0")
		if value > countZeros {
			value = countZeros
			id = number
		}
	}

	countOnes := strings.Count(layers[id], "1")
	countTwos := strings.Count(layers[id], "2")
	ans := countOnes * countTwos
	fmt.Printf("ID: %d, Output: %d\n", id, ans)

	return ans
}

func printLayers(layers []string) {
	for id, layer := range layers {
		countZeros := strings.Count(layer, "0")
		countOnes := strings.Count(layer, "1")
		countTwos := strings.Count(layer, "2")
		fmt.Printf("%03d => %s, %d, %03d, %03d, %03d\n", id, layer, len(layer), countZeros, countOnes, countTwos)
	}
}
