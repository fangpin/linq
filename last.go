package linq

func Last[T comparable](source []T, predictor func(x T) bool) T {
	for i := len(source) - 1; i >= 0; i-- {
		if predictor(source[i]) {
			return source[i]
		}
	}
	panic(ElementNotFound)
}
