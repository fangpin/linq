package linq

func TakeLast[T any](source []T, n int) []T {
	if n < 0 {
		n = 0
	}
	if n > len(source) {
		n = len(source)
	}
	target := make([]T, 0)
	return append(target, source[len(source)-n:]...)
}
