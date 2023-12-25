package day7_part2

func getAnswer(lines []string) int {
	total := 0
	for _, line := range lines {
		listSupernet, listHypernet := getNets(line)
		ok := isTLSSupported(listSupernet, listHypernet)
		if ok {
			total++
		}
	}
	return total
}

func getNets(line string) ([]string, []string) {

	listSupernet := make([]string, 0)
	listHypernet := make([]string, 0)

	addr := ""
	for i, ch := range line {
		switch {
		case ch == ']':
			listHypernet = append(listHypernet, addr)
			addr = ""
			continue
		case ch == '[':
			listSupernet = append(listSupernet, addr)
			addr = ""
			continue
		}
		addr += string(ch)
		if i == len(line)-1 {
			listSupernet = append(listSupernet, addr)
		}
	}
	return listSupernet, listHypernet
}

func isTLSSupported(listSupernet []string, listHypernet []string) bool {

	for _, s := range listHypernet {
		if hasABBA(s) {
			return false
		}
	}

	valid := 0
	for _, s := range listSupernet {
		if hasABBA(s) {
			valid++
		}
	}
	if valid == 0 {
		return false
	}

	return true
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func hasABA(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func hasBAB(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i+1] == s[i] {
			return true
		}
	}
	return false
}

func isSSLSupported(listSupernet []string, listHypernet []string) bool {

	return false
}
