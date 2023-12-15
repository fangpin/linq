package linq

func Keys[T comparable, R any](source map[T]R) []T {
	target := []T{}
	for k, _ := range source {
		target = append(target, k)
	}
	return target
}
