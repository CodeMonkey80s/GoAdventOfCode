package util

//func RemoveAtIndex(s []int, index int) []int {
//	ret := make([]int, 0)
//	ret = append(ret, s[:index]...)
//	return append(ret, s[index+1:]...)
//}

func RemoveAtIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	return append(s[:index], s[index+1:]...)
}
