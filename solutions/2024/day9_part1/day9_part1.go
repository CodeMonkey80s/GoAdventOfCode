package day9_part1

import (
	"GoAdventOfCode/util"
)

const (
	emptySpace = -1
)

func getAnswer(input string) int {

	unpacked := unpackDiskMap(input)
	_, checksum := calculateChecksum(unpacked)

	return checksum
}

func unpackDiskMap(input string) []int {

	var disk []int

	v := 0
	for i := 0; i < len(input); i++ {
		n := util.ConvertByteToInt(input[i])
		switch {
		case i%2 == 0:
			for j := 0; j < n; j++ {
				disk = append(disk, v)
			}
			v++
		default:
			for j := 0; j < n; j++ {
				disk = append(disk, -1)
			}
		}
	}

	return disk
}

func calculateChecksum(disk []int) ([]int, int) {

	a, b := 0, len(disk)-1

	for {
		for a < len(disk) && disk[a] != emptySpace {
			a++
		}

		for b >= 0 && disk[b] == emptySpace {
			b--
		}

		if a >= b {
			break
		}

		disk[a], disk[b] = disk[b], emptySpace
	}

	checksum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == emptySpace {
			continue
		}
		sum := i * disk[i]
		checksum += sum
	}

	return disk, checksum
}
