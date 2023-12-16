package linq

func IntersectBy[T any, K comparable](source []T, keys []K, keyGen func(x T) K) []T {
	set := ToHashSet(keys)
	target := []T{}
	for _, s := range source {
		if _, ok := set[keyGen(s)]; ok {
			target = append(target, s)
		}
	}
	return target
}
