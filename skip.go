package linq

func Skip[T any](source []T, n int) []T {
	if n < 0 || n > len(source) {
		panic(InvalidParam)
	}
	var t []T
	return append(t, source[n:]...)
}
