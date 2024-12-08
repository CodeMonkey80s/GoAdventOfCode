package day8_part2

import (
	"fmt"
	"strings"
)

type antenna struct {
	ch rune
	x  int
	y  int
}

type antinode struct {
	x int
	y int
}

func (a antenna) String() string {
	return fmt.Sprintf("%q %d,%d", a.ch, a.x, a.y)
}

const (
	antennaChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	antennas  []antenna
	antinodes = make(map[antinode]int)

	areaMX = 0
	areaMY = 0
)

func getAnswer(lines []string) int {

	clear(antennas)
	clear(antinodes)

	areaMY = len(lines)
	areaMX = len(lines[0])

	// fmt.Printf("area: %d, %d\n", areaMX, areaMY)

	for y, line := range lines {
		for x, ch := range line {
			switch {
			case ch == '.':
				continue
			case strings.ContainsRune(antennaChars, ch):
				a := antenna{
					ch: ch,
					x:  x,
					y:  y,
				}
				antennas = append(antennas, a)
			default:
				panic(fmt.Sprintf("unknown character %q", ch))
			}
		}
	}

	placeAntinode := func(x int, y int) bool {
		if x < 0 || x >= areaMX || y < 0 || y >= areaMY {
			return false
		}

		anti := antinode{
			x: x,
			y: y,
		}
		antinodes[anti]++

		return true
	}

	placeAntinodes := func(x int, y int, dx int, dy int) {

		for {
			y = y + dy
			x = x + dx

			ok := placeAntinode(x, y)
			if !ok {
				return
			}
		}
	}

	for _, ch := range antennaChars {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {

				a1 := antennas[i]
				a2 := antennas[j]

				placeAntinode(a1.x, a1.y)

				if i == j {
					continue
				}

				if a1.x == a2.x && a1.y == a2.y {
					continue
				}
				if a1.ch != ch || a2.ch != ch {
					continue
				}

				dx := a1.x - a2.x
				dy := a1.y - a2.y

				placeAntinodes(a1.x, a1.y, dx, dy)

			}
		}
	}

	// printArea()

	// for i, a := range antinodes {
	// 	fmt.Printf("antinode: %d, %+v\n", i, a)
	// }

	// fmt.Printf("antennas: %d, antinodes: %d\n", len(antennas), len(antinodes))

	return len(antinodes)
}

func printArea() {
	for y := 0; y < areaMY; y++ {
		for x := 0; x < areaMX; x++ {
			a := antinode{
				x: x,
				y: y,
			}
			_, ok := antinodes[a]
			if ok {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
	fmt.Println()
}
