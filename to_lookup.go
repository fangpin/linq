package linq

func ToLookUp[T, R any, K comparable](source []T, keySelector func(x T) K, elementSelector func(x T) R) map[K]R {
	m := map[K]R{}
	for _, s := range source {
		m[keySelector(s)] = elementSelector(s)
	}
	return m
}
