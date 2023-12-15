package linq

func ToHashSet[T comparable](source []T) map[T]struct{} {
	m := map[T]struct{}{}
	for _, s := range source {
		m[s] = struct{}{}
	}
	return m
}
