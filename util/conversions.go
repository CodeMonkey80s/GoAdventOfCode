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

func ConvertByteToInt(s byte) int {
	val, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return val
}

func ConvertIntToString(n int) string {
	return strconv.Itoa(n)
}
