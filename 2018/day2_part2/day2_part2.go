package day2_part2

import "fmt"

func getAnswer(codes []string) string {
	for _, code1 := range codes {
		for _, code2 := range codes {
			id, ok := compareCodes(code1, code2)
			if ok {
				fmt.Printf("%v = %v\n", code1, code2)
				return id
			}
		}
	}
	return ""
}

func compareCodes(code1 string, code2 string) (string, bool) {
	id := ""
	ok := false
	n := 0
	for i := 0; i < len(code1); i++ {
		if code1[i] != code2[i] {
			n++
			continue
		}
		id += string(code1[i])
	}
	if n == 1 {
		ok = true

	}

	return id, ok
}
