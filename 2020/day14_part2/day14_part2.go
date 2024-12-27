package day14_part2

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"GoAdventOfCode/util"
)

const (
	bitSize = 36
)

var (
	reMask = regexp.MustCompile(`mask = ([0-9a-zA-Z]{36})`)
	reMem  = regexp.MustCompile(`mem\[([0-9]+)\] = ([0-9]+)`)
)

func getAnswer(input []string) int {

	memory := make(map[string]string)

	var mask string
	var bank string
	var value string
	var floating int

	for _, line := range input {

		matches := reMask.FindStringSubmatch(line)
		if len(matches) > 0 {
			mask = matches[1]
			floating = strings.Count(mask, "X")
			// fmt.Printf("mask: %q, floating: %d\n", mask, floating)
			continue
		}

		matches = reMem.FindStringSubmatch(line)
		if len(matches) > 0 {
			bank = util.ConvertIntToBinaryString(util.ConvertStringToInt(matches[1]), bitSize)
			value = matches[2]
			// fmt.Printf("bank: %s, %q, value: %q\n", matches[1], bank, value)
		}

		result := []rune(strings.Clone(bank))
		for i, ch := range mask {
			switch {
			case ch == '0':
				result[i] = rune(bank[i])
			case ch == '1':
				result[i] = '1'
			case ch == 'X':
				result[i] = 'X'
			}
		}

		// fmt.Printf("result: %s\n", string(result))

		for n := range int(math.Pow(2, float64(floating))) {
			addr := []rune(strings.Clone(string(result)))
			bits := util.ConvertIntToBinaryString(n, floating)
			j := 0
			for i, ch := range result {
				switch {
				case ch == '0':
					addr[i] = '0'
				case ch == '1':
					addr[i] = '1'
				case ch == 'X':
					addr[i] = rune(bits[j])
					j++
				}
			}
			// fmt.Printf("n: %d, %s\n", n, string(result))
			// fmt.Printf("n: %d, %s\n", n, string(mask))
			// fmt.Printf("n: %d, %s, %s\n", n, string(addr), value)
			memory[string(addr)] = value
		}

		// fmt.Println()
	}

	sum := 0
	for _, sv := range memory {
		v, err := strconv.ParseInt(sv, 10, 64)
		if err != nil {
			panic(err)
		}
		sum += int(v)
		// fmt.Printf("addr: %s, %d\n", addr, v)
	}

	return sum
}
