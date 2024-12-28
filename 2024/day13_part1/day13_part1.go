package day13_part1

import (
	"regexp"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

type prize struct {
	a point
	b point
	t point
}

func getAnswer(input []string) int {

	tokens := 0
	for i := 0; i < len(input)-4; i += 4 {

		p := loadPrize(input[i : i+4])
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				tx := a*p.a.x + b*p.b.x
				ty := a*p.a.y + b*p.b.y
				if tx == p.t.x && ty == p.t.y {
					cost := a*3 + b
					tokens += cost
				}
			}
		}
	}

	return tokens
}

func loadPrize(input []string) prize {

	var (
		reButtonA = regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)`)
		reButtonB = regexp.MustCompile(`Button B: X\+([0-9]+), Y\+([0-9]+)`)
		rePrize   = regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
	)

	p := prize{}
	for _, line := range input {
		if matches := reButtonA.FindStringSubmatch(line); len(matches) > 0 {
			p.a.x = util.ConvertStringToInt(matches[1])
			p.a.y = util.ConvertStringToInt(matches[2])
		}
		if matches := reButtonB.FindStringSubmatch(line); len(matches) > 0 {
			p.b.x = util.ConvertStringToInt(matches[1])
			p.b.y = util.ConvertStringToInt(matches[2])
		}
		if matches := rePrize.FindStringSubmatch(line); len(matches) > 0 {
			p.t.x = util.ConvertStringToInt(matches[1])
			p.t.y = util.ConvertStringToInt(matches[2])
		}
	}

	return p
}

// func loadPrizeV1(input []string) prize {
//
// 	strA := strings.ReplaceAll(input[0], "Button A:", "")
// 	strA = strings.ReplaceAll(strA, ",", "")
// 	strA = strings.ReplaceAll(strA, "X+", "")
// 	strA = strings.ReplaceAll(strA, "Y+", "")
// 	strA = strings.TrimSpace(strA)
//
// 	strB := strings.ReplaceAll(input[1], "Button B:", "")
// 	strB = strings.ReplaceAll(strB, ",", "")
// 	strB = strings.ReplaceAll(strB, "X+", "")
// 	strB = strings.ReplaceAll(strB, "Y+", "")
// 	strB = strings.TrimSpace(strB)
// 	strPrize := strings.ReplaceAll(input[2], "Prize:", "")
// 	strPrize = strings.ReplaceAll(strPrize, ",", "")
// 	strPrize = strings.ReplaceAll(strPrize, "X=", "")
// 	strPrize = strings.ReplaceAll(strPrize, "Y=", "")
//
// 	buttonAParts := strings.Fields(strA)
// 	buttonBParts := strings.Fields(strB)
// 	prizeParts := strings.Fields(strPrize)
//
// 	p := prize{
// 		a: point{
// 			x: util.ConvertStringToInt(buttonAParts[0]),
// 			y: util.ConvertStringToInt(buttonAParts[1]),
// 		},
// 		b: point{
// 			x: util.ConvertStringToInt(buttonBParts[0]),
// 			y: util.ConvertStringToInt(buttonBParts[1]),
// 		},
// 		t: point{
// 			x: util.ConvertStringToInt(prizeParts[0]),
// 			y: util.ConvertStringToInt(prizeParts[1]),
// 		},
// 	}
//
// 	return p
// }
