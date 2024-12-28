package day13_part2

import (
	"regexp"

	"GoAdventOfCode/util"
)

type point struct {
	x int
	y int
}

type machine struct {
	a point
	b point
	t point
}

func getAnswer(input []string) int {

	const (
		dist = 10_000_000_000_000
	)

	tokens := 0
	for i := 0; i < len(input); i += 4 {

		p := loadMachine(input[i : i+4])

		p.t.x += dist
		p.t.y += dist

		b := (p.t.y*p.a.x - p.t.x*p.a.y) / (p.b.y*p.a.x - p.b.x*p.a.y)
		a := (p.t.x - b*p.b.x) / p.a.x

		if a*p.a.x+b*p.b.x == p.t.x && a*p.a.y+b*p.b.y == p.t.y {
			tokens += a*3 + b
		}

		/*		b = (j*xa - i*ya)/((-ya)*xb + yb*xa);
				a = (i-xb*b)/xa;
				if(a*xa + b*xb == i && a*ya + b*yb == j)
					totalTokens += a*3 + b;
		*/
	}

	return tokens
}

func loadMachine(input []string) machine {

	var (
		reButtonA = regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)`)
		reButtonB = regexp.MustCompile(`Button B: X\+([0-9]+), Y\+([0-9]+)`)
		rePrize   = regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
	)

	p := machine{}
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
