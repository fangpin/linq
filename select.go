package linq

func Select[T, R any](source []T, f func(x T) R) []R {
	target := make([]R, len(source))
	for i, s := range source {
		target[i] = f(s)
	}
	return target
}
