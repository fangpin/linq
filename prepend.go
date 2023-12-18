package linq

func Prepend[T any](source []T, target ...T) []T {
	t := Reverse(target)
	return append(t, source...)
}
