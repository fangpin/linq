package linq

func ToDictionary[T any, K comparable](source []T, keyGen func(x T) K) map[K]T {
	m := map[K]T{}
	for _, s := range source {
		m[keyGen(s)] = s
	}
	return m
}
