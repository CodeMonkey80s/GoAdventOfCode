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

func ConvertByteToInt(s byte) int {
	val, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return val
}
