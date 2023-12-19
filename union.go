package linq

func Union[T comparable](left, right []T) []T {
	m1 := ToHashSet(left)
	m2 := ToHashSet(right)
	for k := range m2 {
		m1[k] = struct{}{}
	}
	return Keys(m1)
}
