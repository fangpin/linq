package linq

func SkipWhile[T any](source []T, skipCond func(x T) bool) []T {
	return Where(source, func(x T) bool { return !skipCond(x) })
}
