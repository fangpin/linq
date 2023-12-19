package linq

func Single[T comparable](source []T, pred func(x T) bool) T {
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
	return result
}
