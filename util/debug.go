package util

import "fmt"

func PrintSlice[T any](name string, slice []T) {
	fmt.Println("*** " + name + " ***")
	for _, v := range slice {
		fmt.Printf("\"%v\"\n", v)
	}
}
