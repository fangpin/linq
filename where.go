package linq

func Where[T any](source []T, f func(x T) bool) []T {
	target := make([]T, 0)
	for _, s := range source {
		if f(s) {
			target = append(target, s)
		}
	}
	return target
}
