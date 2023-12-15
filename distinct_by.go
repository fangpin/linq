package linq

func DistinctBy[T any, R comparable](source []T, f func(x T) R) []T {
	m := map[R]T{}
	for _, s := range source {
		key := f(s)
		if _, ok := m[key]; !ok {
			m[key] = s
		}
	}
	return Values(m)
}
