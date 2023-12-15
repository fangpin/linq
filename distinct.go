package linq

func Distinct[T comparable](source []T) []T {
	set := map[T]struct{}{}
	for _, s := range source {
		set[s] = struct{}{}
	}
	return Keys(set)
}
