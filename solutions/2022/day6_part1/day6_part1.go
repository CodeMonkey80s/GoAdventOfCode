package day6_part1

func startOfPacket(input string) int {
	a := 0
	b := 4
	for i := 0; i < len(input)-4; i++ {
		ok := isMarker(input[a:b])
		if ok {
			return a + 4
		}
		a++
		b++
	}

	return 0
}

func isMarker(input string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(input); i++ {
		ch := input[i]
		m[ch]++
	}
	if len(m) == 4 {
		return true
	}
	return false
}
