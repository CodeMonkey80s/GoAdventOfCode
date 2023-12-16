package day1_part2

func getAnswer(input string) int {
	level := 0
	for i, char := range input {
		switch char {
		case '(':
			level++
		case ')':
			level--
		}
		if level == -1 {
			return i + 1
		}
	}
	return 0
}
