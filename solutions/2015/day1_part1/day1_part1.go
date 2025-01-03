package day1_part1

func getAnswer(input string) int {
	level := 0
	for _, char := range input {
		switch char {
		case '(':
			level++
		case ')':
			level--
		}
	}
	return level
}
