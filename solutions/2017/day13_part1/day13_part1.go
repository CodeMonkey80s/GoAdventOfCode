package day13_part1

import (
	"regexp"

	"GoAdventOfCode/util"
)

type packet struct {
	depth int
}

type scanner struct {
	depth    int
	maxRange int
	curRange int
	d        int
}

func getAnswer(input []string) int {

	var scanners []scanner
	var p packet

	maxPicoseconds := 0
	for _, line := range input {

		s := loadScanner(line)
		scanners = append(scanners, s)
		maxPicoseconds = max(maxPicoseconds, s.depth)
	}

	severity := 0
	picoseconds := 0
	for picoseconds <= maxPicoseconds {
		for _, s := range scanners {
			if p.depth == s.depth && s.curRange == 0 {
				severity += p.depth * s.maxRange
			}
		}
		for i := range scanners {
			scanners[i].curRange += scanners[i].d
			if scanners[i].curRange >= scanners[i].maxRange-1 || scanners[i].curRange == 0 {
				scanners[i].d = -scanners[i].d
			}
		}
		p.depth++
		picoseconds++
	}

	return severity
}

func loadScanner(line string) scanner {

	var (
		reScanner = regexp.MustCompile(`([0-9]+): ([0-9]+)`)
	)

	s := scanner{}
	if matches := reScanner.FindStringSubmatch(line); len(matches) > 0 {
		s.depth = util.ConvertStringToInt(matches[1])
		s.maxRange = util.ConvertStringToInt(matches[2])
		s.d = 1
	}

	return s
}
