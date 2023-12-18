package linq

func Reverse[T any](source []T) (target []T) {
	target = make([]T, len(source))
	for i, l := 0, len(source); i < l; i++ {
		target[i] = source[l-i-1]
	}
	return target
}
