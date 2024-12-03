package day3_part2

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	regexExpression = `(do\(\)|don't\(\))|(mul\(([\d]{1,3}),([\d]{1,3})\))`

	strDo   = "do()"
	strDont = "don't()"
)

func getAnswer(input []string) int {

	text := strings.Join(input, "")
	text = strings.ReplaceAll(text, "\n", "")

	sum := 0
	re := regexp.MustCompile(regexExpression)
	matches := re.FindAllStringSubmatch(text, -1)

	do := true
	for _, match := range matches {

		if slices.Contains(match, strDont) {
			do = false
			continue
		}

		if slices.Contains(match, strDo) {
			do = true
			continue
		}

		if do {
			num1, _ := strconv.Atoi(match[len(match)-2])
			num2, _ := strconv.Atoi(match[len(match)-1])
			sum += num1 * num2
		}
	}

	return sum
}
