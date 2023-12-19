package linq

func UnionBy[T any, K comparable](left, right []T, keySelector func(x T) K) []T {
	m := map[K]T{}
	for _, s := range left {
		m[keySelector(s)] = s
	}
	for _, s := range right {
		m[keySelector(s)] = s
	}
	return Values(m)
}
