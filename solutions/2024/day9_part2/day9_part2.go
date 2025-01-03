package day9_part2

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

func findOnDiskID(disk []int, id int) (int, int, bool) {
	a := -1
	b := -1
	// fmt.Printf("searching for id %d\n", id)
	for i := len(disk) - 1; i >= 0; i-- {
		if disk[i] == id && b == -1 {
			b = i
			for j := i; j >= 0; j-- {
				a = j
				if disk[j] != id {
					a++
					b++
					return a, b, true
				}
			}
		}
	}

	return a, b, false
}

func findOnDiskEmptySpace(disk []int, idx int, size int) (int, int, bool) {
	a := -1
	b := -1
	start := 0
	// fmt.Printf("searching for empty space of size %d\n", size)
outer:
	for {
		if start > len(disk) || start >= idx {
			break
		}
		for i := start; i < len(disk); i++ {
			if disk[i] == emptySpace {
				a = i
				for j := i; j < len(disk); j++ {
					if disk[j] != emptySpace {
						b = j
						// fmt.Printf("a: %d, b: %d, idx: %d, b-a: %d, size: %d\n", a, b, idx, b-a, size)
						if b-a >= size && b < idx {
							return a, b, true
						}

						if b > idx {
							// fmt.Printf("space not found!\n")
							return a, b, false
						}

						start = b
						continue outer
					}
				}
			}
		}
		start++
	}

	// fmt.Printf("space not found!\n")
	return a, b, false
}

func moveOnDiskID(disk []int, a1 int, a2 int, b1 int, b2 int) []int {
	for i := 0; i < b2-b1; i++ {
		disk[a1+i], disk[b1+i] = disk[b1+i], emptySpace
	}

	// fmt.Printf("disk: %#v\n", disk)

	return disk
}

func calculateChecksum(disk []int) ([]int, int) {

	id := disk[len(disk)-1]

	for {

		// fmt.Printf("id %d\n", id)
		b1, b2, ok1 := findOnDiskID(disk, id)
		if ok1 {
			size := b2 - b1
			// fmt.Printf("id: %d, size: %d, layout: %#v\n", id, size, disk[b1:b2])
			a1, a2, ok2 := findOnDiskEmptySpace(disk, b2, size)
			if ok2 {
				// fmt.Printf("space found at %d %d\n", a1, a2)
				disk = moveOnDiskID(disk, a1, a2, b1, b2)
				// fmt.Printf("moved id %d\n", id)
			}
		}

		// fmt.Println()
		id--
		if id < 0 {
			break
		}
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
