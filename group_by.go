package linq

func GroupBy[T any, R comparable, S any](source []T, keyGen func(x T) R, groupSelector func([]T) S) map[R]S {
	group := map[R][]T{}
	for _, s := range source {
		key := keyGen(s)
		elements := group[key]
		elements = append(elements, s)
		group[key] = elements
	}
	target := map[R]S{}
	for k, v := range group {
		target[k] = groupSelector(v)
	}
	return target
}
