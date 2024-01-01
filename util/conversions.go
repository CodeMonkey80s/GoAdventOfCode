package util

import (
	"strconv"
)

func ConvertStringToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}

func ConvertStringToUint16(s string) uint16 {
	val, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0
	}
	return uint16(val)
}

func ConvertBinaryStringToInt(s string) int {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0
	}
	return int(val)
}

func ConvertIntToBinaryString(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func ConvertByteToInt(s byte) int {
	val, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return val
}

func ConvertRuneToInt(s rune) int {
	val, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return val
}

func ConvertIntToString(n int) string {
	return strconv.Itoa(n)
}

func IsNumber(s string) bool {
	for _, char := range s {
		if '0' >= char || char >= '9' {
			return false
		}
	}
	return true
}

func ConvertIntToLetters(n int) string {
	result := ""
	for n > 0 {
		remainder := (n - 1) % 26
		result = string(rune('A'+remainder)) + result
		n = (n - 1) / 26
	}
	if n < 26 {
		result = string(rune('A'+n)) + result
	}
	return result
}
