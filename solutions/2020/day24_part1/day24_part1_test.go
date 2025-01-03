package day24_part1

import (
	"fmt"
	"testing"

	"GoAdventOfCode/util"
)

func Test_getAnswer(t *testing.T) {
	var testCases = []struct {
		Input  []string
		Output int
	}{
		{
			Input: []string{
				"nwwswee",
			},
			Output: 1,
		},
		{
			Input: []string{
				"esenee",
			},
			Output: 1,
		},
		{
			Input: []string{
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
			},
			Output: 10,
		},
	}
	lines := util.LoadInputFile("../inputs/day24_input.txt")
	testCase := []struct {
		Input  []string
		Output int
	}{
		{
			Input:  lines,
			Output: 346,
		},
	}
	testCases = append(testCases, testCase...)
	for _, tc := range testCases {
		label := fmt.Sprintf("%v\n", "Puzzle Input")
		t.Run(label, func(t *testing.T) {
			output := getAnswer(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be '%d' but we got '%d'", tc.Output, output)
			}
		})
	}
}
