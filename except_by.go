package linq

func ExceptBy[T any, R comparable](source []T, keys []R, f func(x T) R) []T {
	set := ToHashSet(keys)
	target := []T{}
	for _, s := range source {
		if _, ok := set[f(s)]; !ok {
			target = append(target, s)
		}
	}
	return target
}
