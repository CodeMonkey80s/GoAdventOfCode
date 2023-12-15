package day4_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

type Board struct {
	Layout []string
}

func finalScore(lines []string) int {

	numbers := getNumbers(lines[0])
	boards := getBoards(lines[2:])
	score := play(numbers, boards)

	return score
}

func getNumbers(line string) []string {
	numbers := make([]string, 0)
	parts := strings.Split(line, ",")
	for _, number := range parts {
		numbers = append(numbers, number)
	}
	return numbers
}

func getBoards(lines []string) []*Board {
	boards := make([]*Board, 0)
	for i := 0; i < len(lines); i += 6 {
		b := Board{
			Layout: make([]string, 0),
		}
		b.Layout = append(b.Layout, lines[i:i+5]...)
		boards = append(boards, &b)
	}
	return boards
}

type Wins struct {
	Board      *Board
	LastNumber int
	N          int
}

func play(numbers []string, boards []*Board) int {
	wins := make([]*Wins, 0)
	for _, number := range numbers {
		for b, board := range boards {
			for i, layout := range board.Layout {
				line := ""
				parts := strings.Fields(layout)
				for j, n := range parts {
					if n == number {
						line += "XX"
					} else {
						l := len(n)
						if l == 1 {
							line += " "
						}
						line += n
					}
					if j < len(parts)-1 {
						line += " "
					}
				}
				board.Layout[i] = line
				if isBingo(board) {
					w := Wins{
						Board:      board,
						LastNumber: util.ConvertStringToInt(number),
						N:          b,
					}
					found := false
					for _, w := range wins {
						if w.Board == board {
							found = true
						}
					}
					if !found {
						wins = append(wins, &w)
					}
					if len(wins) >= len(boards) {
						last := wins[len(wins)-1]
						return last.LastNumber * sumNumberForBoard(last.Board)
					}
				}

			}
		}
	}
	return 0
}

func isBingo(board *Board) bool {
	for _, layout := range board.Layout {
		n := strings.Count(layout, "XX")
		if n == 5 {
			return true
		}
	}
	for x := 0; x < 11; x += 3 {
		c := 0
		for y := 0; y < 5; y++ {
			if board.Layout[y][x:x+2] == "XX" {
				c++
			}
		}
		if c == 5 {
			return true
		}
	}
	return false
}

func sumNumberForBoard(board *Board) int {
	sum := 0
	for _, layout := range board.Layout {
		parts := strings.Fields(layout)
		for _, number := range parts {
			if number == "XX" {
				continue
			}
			sum += util.ConvertStringToInt(number)
		}
	}
	return sum
}
