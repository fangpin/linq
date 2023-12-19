package linq

func SingleOrDefault[T comparable](source []T, pred func(x T) bool) T {
	var result T
	found := false
	for _, s := range source {
		if pred(s) {
			if found {
				panic("multiple items found")
			}
			found = true
			result = s
		}
	}
	if !found {
		result = zero[T]{}.inner
	}
	return result
}
