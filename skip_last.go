package linq

func SkipLast[T any](source []T, count int) []T {
	if count < 0 {
		panic("negative skip account")
	}
	if len(source) < count {
		panic("skip account exceeds the source length")
	}
	var t []T
	return append(t, source[0:len(source)-count]...)
}
