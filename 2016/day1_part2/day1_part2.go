package day1_part2

import (
	"math"
	"strings"

	"GoAdventOfCode/util"
)

var visited = make(map[Point]int)

type Step struct {
	Dir    rune
	Length int
}

type Point struct {
	X int
	Y int
}

type Player struct {
	Facing     rune
	Position   Point
	HQPosition Point
	Found      bool
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
		sy := p.Position.Y - 1
		p.Position.Y += s.Length * dy
		for y := sy; y >= p.Position.Y; y-- {
			v := Point{
				X: p.Position.X,
				Y: y,
			}
			visited[v]++
			//fmt.Printf("N: %v, %d\n", v, visited[v])
			if visited[v] == 2 && !p.Found {
				p.HQPosition = v
				p.Found = true
				return
			}
		}
	case p.Facing == 'S':
		dx = 0
		dy = 1
		sy := p.Position.Y + 1
		p.Position.Y += s.Length * dy
		for y := sy; y <= p.Position.Y; y++ {
			v := Point{
				X: p.Position.X,
				Y: y,
			}
			visited[v]++
			//fmt.Printf("S: %v, %d\n", v, visited[v])
			if visited[v] == 2 && !p.Found {
				p.HQPosition = v
				p.Found = true
				return
			}
		}
	case p.Facing == 'W':
		dx = -1
		dy = 0
		sx := p.Position.X - 1
		p.Position.X += s.Length * dx
		for x := sx; x >= p.Position.X; x-- {
			v := Point{
				X: x,
				Y: p.Position.Y,
			}
			visited[v]++
			//fmt.Printf("W: %v, %d\n", v, visited[v])
			if visited[v] == 2 && !p.Found {
				p.HQPosition = v
				p.Found = true
				return
			}
		}
	case p.Facing == 'E':
		dx = 1
		dy = 0
		sx := p.Position.X + 1
		p.Position.X += s.Length * dx
		for x := sx; x <= p.Position.X; x++ {
			v := Point{
				X: x,
				Y: p.Position.Y,
			}
			visited[v]++
			//fmt.Printf("E: %v, %d\n", v, visited[v])
			if visited[v] == 2 && !p.Found {
				p.HQPosition = v
				p.Found = true
				return
			}
		}
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
	visited = make(map[Point]int)
	start := Point{}
	p := &Player{
		Facing: 'N',
		Position: Point{
			X: 0,
			Y: 0,
		},
		HQPosition: Point{
			X: 0,
			Y: 0,
		},
		Found: false,
	}
	steps := getSteps(input)
	for _, step := range steps {
		p.Walk(step)
		if p.Found {
			d := calculateDistance(start, p.HQPosition)
			return d
		}
	}
	return -1
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
	return sum
}
