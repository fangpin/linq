package linq

func First[T any](source []T, f func(x T) bool) T {
	for _, s := range source {
		if f(s) {
			return s
		}
	}
	panic("not found satisfied item")
}
