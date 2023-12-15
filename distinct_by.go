package linq

func DistinctBy[T any, R comparable](source []T, f func(x T) R) []T {
	m := map[R]T{}
	for _, s := range source {
		key := f(s)
		m[key] = s
	}
	return Values(m)
}
