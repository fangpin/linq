package linq

func Take[T any](source []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(source) {
		n = len(source)
	}
	target := make([]T, n)
	for i := 0; i < n; i++ {
		target[i] = source[i]
	}
	return target
}
