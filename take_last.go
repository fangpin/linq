package linq

func TakeLast[T any](source []T, n int) []T {
	if n < 0 {
		panic(InvalidParam)
	}
	if n > len(source) {
		panic(InvalidParam)
	}
	target := make([]T, 0)
	return append(target, source[len(source)-n:]...)
}
