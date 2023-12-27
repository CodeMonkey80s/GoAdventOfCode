package util

import "fmt"

func PrintSlice[T any](name string, slice []T) {
	fmt.Println("*** " + name + " ***")
	for _, v := range slice {
		fmt.Printf("\"%v\"\n", v)
	}
}

func PrintMap[T1 comparable, T2 any](name string, m map[T1]T2) {
	fmt.Println("*** " + name + " ***")
	for k, v := range m {
		fmt.Printf("%v => \"%+v\"\n", k, v)
	}
}
