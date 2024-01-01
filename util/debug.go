package util

import (
	"cmp"
	"fmt"
	"slices"
)

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

func PrintOrderedMap[T1 cmp.Ordered, T2 any](name string, m map[T1]T2) {
	fmt.Println("*** " + name + " ***")
	keys := make([]T1, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	fmt.Println(keys)
	for _, v := range keys {
		fmt.Printf("%s => \"%+v\"\n", v, m[v])
	}
}
