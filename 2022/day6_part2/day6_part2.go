package day6_part2

const (
	length = 14
)

func startOfPacket(input string) int {
	a := 0
	b := length
	for i := 0; i < len(input)-length; i++ {
		ok := isMessage(input[a:b])
		if ok {
			return a + length
		}
		a++
		b++
	}

	return 0
}

func isMessage(input string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		ch := input[i]
		m[ch]++
	}
	if len(m) == length {
		return true
	}
	return false
}
