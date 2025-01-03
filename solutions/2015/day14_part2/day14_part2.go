package day14_part2

import (
	"strings"

	"GoAdventOfCode/util"
)

type State string

const (
	stateFlying  = "flying"
	stateResting = "resting"
)

type Deer struct {
	Name        string
	Speed       int
	TimeAtMove  int
	TimeAtRest  int
	TimeAtState int
	Travelled   int
	Points      int
	State       State
}

func getAnswer(lines []string, flyTime int) int {

	deers := make([]*Deer, 0)
	for _, line := range lines {
		deer := getDeer(line)
		deers = append(deers, &deer)
	}

	fly(deers, flyTime)

	points := 0
	for _, deer := range deers {
		points = max(points, deer.Points)
	}

	return points
}

func getDeer(line string) Deer {
	parts := strings.Fields(line)

	deer := Deer{
		Name:       parts[0],
		Speed:      util.ConvertStringToInt(parts[3]),
		TimeAtMove: util.ConvertStringToInt(parts[6]),
		TimeAtRest: util.ConvertStringToInt(parts[13]),
		State:      stateFlying,
	}

	return deer
}

func fly(deers []*Deer, flyTime int) {
	for i := 0; i < flyTime; i++ {
		for _, deer := range deers {
			switch {
			case deer.State == stateFlying:
				deer.Travelled += deer.Speed
				deer.TimeAtState++
				if deer.TimeAtState >= deer.TimeAtMove {
					deer.State = stateResting
					deer.TimeAtState = 0
				}
			case deer.State == stateResting:
				deer.TimeAtState++
				if deer.TimeAtState >= deer.TimeAtRest {
					deer.State = stateFlying
					deer.TimeAtState = 0
				}
			}
		}
		distance := 0
		for _, deer := range deers {
			distance = max(distance, deer.Travelled)
		}
		for _, deer := range deers {
			if deer.Travelled == distance {
				deer.Points++
			}
		}
	}
}
