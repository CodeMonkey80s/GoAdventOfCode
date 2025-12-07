package day5_part2

import (
	"GoAdventOfCode/util"
	"slices"
	"strings"
)

func getAnswerForPart2(lines []string) int {

	ranges := loadRanges(lines)

	slices.SortFunc(ranges, func(a, b valueRange) int {
		return a.a - b.a
	})

	var separate []valueRange

loop:
	for i := 0; i < len(ranges)-1; i++ {
		a := ranges[i].a
		b := ranges[i].b
		c := ranges[i+1].a
		d := ranges[i+1].b
		switch {
		// when ranges are separate
		case b < c || d < a:
			separate = append(separate, ranges[i])
			ranges = slices.Delete(ranges, i, i+1)
			goto loop
		// when ranges overlap
		case a <= d && c <= b:
			ranges[i].a = min(a, c)
			ranges[i].b = max(b, d)
			ranges = slices.Delete(ranges, i+1, i+2)
			goto loop
		}
	}

	var numbers int
	for _, r := range separate {
		numbers += r.b - r.a + 1
	}

	for _, r := range ranges {
		numbers += r.b - r.a + 1
	}

	return numbers
}

type valueRange struct {
	a int
	b int
}

func loadRanges(lines []string) []valueRange {
	var ranges []valueRange
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		if strings.IndexByte(line, '-') != -1 {
			parts := strings.Split(line, "-")
			r := valueRange{
				a: util.ConvertStringToInt(parts[0]),
				b: util.ConvertStringToInt(parts[1]),
			}
			ranges = append(ranges, r)
		}
	}
	return ranges
}
