package day7_part1

func getAnswer(lines []string) int {
	total := 0
	for _, line := range lines {
		ok := isTLSSupported(line)
		if ok {
			total++
		}
	}
	return total
}

func isTLSSupported(line string) bool {

	listSegments := make([]string, 0)
	listHypernet := make([]string, 0)

	addr := ""
	for i, ch := range line {
		switch {
		case ch == ']':
			listHypernet = append(listHypernet, addr)
			addr = ""
			continue
		case ch == '[':
			listSegments = append(listSegments, addr)
			addr = ""
			continue
		}
		addr += string(ch)
		if i == len(line)-1 {
			listSegments = append(listSegments, addr)
		}
	}

	for _, s := range listHypernet {
		if isABBA(s) {
			return false
		}
	}

	valid := 0
	for _, s := range listSegments {
		if isABBA(s) {
			valid++
		}
	}
	if valid == 0 {
		return false
	}

	return true
}

func isABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}
