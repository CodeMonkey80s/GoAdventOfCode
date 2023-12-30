package day8_part2

import (
	"fmt"
)

func getAnswer(input string, w int, h int) string {

	layers := make([]string, 0)

	n := w * h
	for i, d := range input {
		if i%n == 0 {
			layers = append(layers, "")
		}
		layers[i/n] += string(d)
	}

	//printLayers(layers)
	printDecodedLayers(layers, n, w, h)

	return layers[0]
}

func printLayers(layers []string) {
	for id, layer := range layers {
		fmt.Printf("Layer %d: %s\n", id, layer)
	}
}

func printDecodedLayers(layers []string, n int, w int, h int) {
	// decode the layers and print the result in a grid
	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)
	}

	// Iterate through each pixel in the image
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			// Find the color of the pixel by iterating through the layers
			for _, layer := range layers {
				pixel := layer[y*w+x]
				if pixel != '2' {
					// Set the color of the pixel in the grid
					grid[y][x] = int(pixel - '0')
					break
				}
			}
		}
	}

	// Print the decoded image
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
