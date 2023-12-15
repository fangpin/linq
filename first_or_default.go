package linq

func FirstOrDefault[T any](source []T, pred func(x T) bool) T {
	for _, s := range source {
		if pred(s) {
			return s
		}
	}
	return zero[T]{}.inner
}
