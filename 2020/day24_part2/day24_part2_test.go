package day24_part2

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

var input = []string{
	"sesenwnenenewseeswwswswwnenewsewsw",
	"neeenesenwnwwswnenewnwwsewnenwseswesw",
	"seswneswswsenwwnwse",
	"nwnwneseeswswnenewneswwnewseswneseene",
	"swweswneswnenwsewnwneneseenw",
	"eesenwseswswnenwswnwnwsewwnwsene",
	"sewnenenenesenwsewnenwwwse",
	"wenwwweseeeweswwwnwwe",
	"wsweesenenewnwwnwsenewsenwwsesesenwne",
	"neeswseenwwswnwswswnw",
	"nenwswwsewswnenenewsenwsenwnesesenew",
	"enewnwewneswsewnwswenweswnenwsenwsw",
	"sweneswneswneneenwnewenewwneswswnese",
	"swwesenesewenwneswnwwneseswwne",
	"enesenwswwswneneswsenwnewswseenwsese",
	"wnwnesenesenenwwnenwsewesewsesesew",
	"nenewswnwewswnenesenwnesewesw",
	"eneswnwswnwsenenwnwnwwseeswneewsenese",
	"neswnwewnwnwseenwseesewsenwsweewe",
	"wseweeenwnesenwwwswnew",
}

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: input,
			InputB: 1,
			Output: 15,
		},
		{
			InputA: input,
			InputB: 2,
			Output: 12,
		},
		{
			InputA: input,
			InputB: 3,
			Output: 25,
		},
		{
			InputA: input,
			InputB: 4,
			Output: 14,
		},
		{
			InputA: input,
			InputB: 100,
			Output: 2208,
		},
	}
	lines := util.LoadInputFile("../inputs/day24_input.txt")
	testCase := []struct {
		InputA []string
		InputB int
		Output int
	}{
		{
			InputA: lines,
			InputB: 100,
			Output: 0,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
