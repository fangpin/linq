package linq

func SelectMany[T, R any](source []T, manySelector func(x T) []R) []R {
	target := make([]R, 0)
	for _, s := range source {
		target = append(target, manySelector(s)...)
	}
	return target
}
