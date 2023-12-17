package day1_part1

import (
	"fmt"
	"math"
	"strings"

	"GoAdventOfCode/util"
)

type Step struct {
	Dir    rune
	Length int
}

func (step Step) String() string {
	return fmt.Sprintf("Dir: '%c', Length: %v", step.Dir, step.Length)
}

type Point struct {
	X int
	Y int
}

type Player struct {
	Facing   rune
	Position Point
}

func (p *Player) String() string {
	return fmt.Sprintf("Facing: '%c', X: %d Y: %d", p.Facing, p.Position.X, p.Position.Y)
}

//   N
// W   E
//   S

func (p *Player) Walk(s Step) {
	dx := 0
	dy := 0
	p.updateFacing(s.Dir)
	switch {
	case p.Facing == 'N':
		dx = 0
		dy = -1
	case p.Facing == 'S':
		dx = 0
		dy = 1
	case p.Facing == 'W':
		dx = -1
		dy = 0
	case p.Facing == 'E':
		dx = 1
		dy = 0
	}
	if dx != 0 {
		p.Position.X += s.Length * dx
	}
	if dy != -0 {
		p.Position.Y += s.Length * dy
	}
}

func (p *Player) updateFacing(dir rune) {
	switch {
	case p.Facing == 'N' && dir == 'L':
		p.Facing = 'W'
	case p.Facing == 'N' && dir == 'R':
		p.Facing = 'E'
	case p.Facing == 'S' && dir == 'L':
		p.Facing = 'E'
	case p.Facing == 'S' && dir == 'R':
		p.Facing = 'W'
	case p.Facing == 'W' && dir == 'L':
		p.Facing = 'S'
	case p.Facing == 'W' && dir == 'R':
		p.Facing = 'N'
	case p.Facing == 'E' && dir == 'L':
		p.Facing = 'N'
	case p.Facing == 'E' && dir == 'R':
		p.Facing = 'S'
	}
}

func getAnswer(input string) int {
	start := Point{}
	p := &Player{
		Facing: 'N',
		Position: Point{
			X: 0,
			Y: 0,
		},
	}
	steps := getSteps(input)
	for _, step := range steps {
		p.Walk(step)
		//fmt.Printf("i: %d -> %+v\n", i, p)
	}

	return calculateDistance(start, p.Position)
}

func getSteps(input string) []Step {
	steps := make([]Step, 0)
	list := strings.Split(input, ", ")
	for _, v := range list {
		step := Step{
			Dir:    rune(v[0]),
			Length: util.ConvertStringToInt(v[1:]),
		}
		steps = append(steps, step)
	}

	return steps
}

func calculateDistance(p1 Point, p2 Point) int {
	dx := int(math.Abs(float64(p1.X - p2.X)))
	dy := int(math.Abs(float64(p1.Y - p2.Y)))
	sum := dx + dy
	//fmt.Printf("dx: %d, dy: %d, sum: %d\n", dx, dy, sum)
	return sum
}
