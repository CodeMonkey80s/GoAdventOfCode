package util

// PermString calls f with each permutation of a.
func PermString(a []string, f func([]string)) {
	permString(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func permString(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	permString(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permString(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
