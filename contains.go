package linq

func Contains[T comparable](source []T, key T) bool {
	return Any(source, func(x T) bool {
		return x == key
	})
}
