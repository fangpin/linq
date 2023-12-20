package linq

func SkipLast[T any](source []T, count int) []T {
	if count < 0 {
		panic(InvalidParam)
	}
	if len(source) < count {
		panic(InvalidParam)
	}
	var t []T
	return append(t, source[0:len(source)-count]...)
}
