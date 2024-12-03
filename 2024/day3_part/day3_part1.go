package day3_part

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	regexExpression = `mul\(([\d]{1,3}),([\d]{1,3})\)`
)

func getAnswer(input []string) int {

	text := strings.Join(input, "")
	text = strings.ReplaceAll(text, "\n", "")

	sum := 0
	re := regexp.MustCompile(regexExpression)
	matches := re.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += num1 * num2
	}

	return sum
}
