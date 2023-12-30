package day7_part2

func getAnswer(lines []string) int {
	total := 0
	for _, line := range lines {
		listSupernet, listHypernet := getNets(line)
		ok := isSSLSupported(listSupernet, listHypernet)
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

func getABAs(s string) ([]string, bool) {
	abas := make([]string, 0)
	found := false
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			aba := string(s[i]) + string(s[i+1]) + string(s[i])
			abas = append(abas, aba)
			found = true
		}
	}
	return abas, found
}

func hasBAB(s string, aba string) bool {
	bab := string(aba[1]) + string(aba[0]) + string(aba[1])
	for i := 0; i < len(s)-2; i++ {
		if s[i] == bab[0] && s[i+1] == bab[1] && s[i+2] == bab[2] {
			return true
		}
	}
	return false
}

func isSSLSupported(listSupernet []string, listHypernet []string) bool {
	for _, sn := range listSupernet {
		abas, ok := getABAs(sn)
		if ok {
			for _, aba := range abas {
				for _, hn := range listHypernet {
					if hasBAB(hn, aba) {
						return true
					}
				}
			}
		}
	}

	return false
}
