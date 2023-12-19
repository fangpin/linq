package linq

func TakeWhile[T any](source []T, takeCond func(x T) bool) []T {
	return Where(source, takeCond)
}
