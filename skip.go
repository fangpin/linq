package linq

func Skip[T any](source []T, n int) []T {
	if n < 0 {
		panic("negative skip number")
	}
	var t []T
	return append(t, source[n:]...)
}
