package day4_part2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"GoAdventOfCode/util"
)

func getAnswer(lines []string) int {
	ans := 0
	sb := strings.Builder{}
	for i, line := range lines {
		_, err := sb.WriteString(line)
		if err != nil {
			panic(err)
		}
		sb.WriteByte(' ')
		if line == "" || i == len(lines)-1 {
			pass := sb.String()
			pass = strings.TrimSpace(pass)
			if isValid(pass) {
				ans++
			}
			sb.Reset()
		}
	}

	return ans
}

func isValid(s string) bool {
	parts := strings.Fields(s)
	m := make(map[string]bool)
	for _, part := range parts {
		field := part[0:3]
		m[field] = true
	}

	a := 0
	_, ok := m["cid"]
	if ok {
		a = -1
	} else {
		a = 0
	}

	fmt.Printf("s: %q, len: %d\n", s, len(m)+a)
	if len(m)+a <= 6 {
		fmt.Println("***** invalid number of fields", len(m), ok)
		return false
	}

	for _, part := range parts {
		field := part[0:3]
		switch field {
		case "byr":
			p := strings.Split(part, ":")
			year := util.ConvertStringToInt(p[1])
			if 1920 > year || year > 2002 || len(p[1]) != 4 {
				fmt.Println("***** invalid byr", year)
				return false
			}
		case "iyr":
			p := strings.Split(part, ":")
			year := util.ConvertStringToInt(p[1])
			if 2010 > year || year > 2020 || len(p[1]) != 4 {
				fmt.Println("***** invalid iyr", year)
				return false
			}
		case "eyr":
			p := strings.Split(part, ":")
			year := util.ConvertStringToInt(p[1])
			if 2020 > year || year > 2030 || len(p[1]) != 4 {
				fmt.Println("***** invalid eyr", year)
				return false
			}
		case "hgt":
			p := strings.Split(part, ":")
			height := util.ConvertStringToInt(p[1][0 : len(p[1])-2])
			unit := p[1][len(p[1])-2:]
			if unit != "cm" && unit != "in" {
				fmt.Println("***** invalid hgt", unit)
				return false
			}
			if unit == "cm" {
				if 150 > height || height > 193 {
					fmt.Println("***** invalid hgt", height, unit)
					return false
				}
			}
			if unit == "in" {
				if 59 > height || height > 76 {
					fmt.Println("***** invalid hgt", height, unit)
					return false
				}
			}
		case "hcl":
			p := strings.Split(part, ":")
			if !strings.HasPrefix(p[1], "#") {
				fmt.Println("***** invalid hcl", p[1])
				return false
			}
			if len(p[1]) != 7 {
				fmt.Println("***** invalid hcl", p[1])
				return false
			}
			n, err := strconv.ParseUint(p[1][1:], 16, 64)
			if err != nil || n == 0 {
				fmt.Println("***** invalid hcl", p[1])
				return false
			}
		case "ecl":
			p := strings.Split(part, ":")
			eyes := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			if !slices.Contains(eyes, p[1]) {
				fmt.Println("***** invalid ecl", p[1])
				return false
			}
		case "pid":
			p := strings.Split(part, ":")
			if len(p[1]) != 9 {
				fmt.Println("***** invalid pid", p[1])
				return false
			}
		}
	}

	return true
}
