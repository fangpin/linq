package linq

func TakeLast[T any](source []T, count int) []T {
	return Take(Reverse(source), count)
}
