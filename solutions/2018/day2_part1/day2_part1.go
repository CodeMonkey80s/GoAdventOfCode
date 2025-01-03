package day2_part1

func getAnswer(codes []string) int {
	two := 0
	three := 0
	for _, code := range codes {
		freq := make(map[rune]int)
		for _, char := range code {
			freq[char]++
		}
		ok2 := false
		ok3 := false
		for _, v := range freq {
			if v == 2 {
				ok2 = true
			}
			if v == 3 {
				ok3 = true
			}
		}
		if ok2 && ok3 {
			two++
			three++
			continue
		}
		if ok2 {
			two++
			continue
		}
		if ok3 {
			three++
			continue
		}
	}
	return two * three
}
