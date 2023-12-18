package linq

func Repeat[T any](element T, times uint) []T {
	ret := make([]T, times)
	for i := uint(0); i < times; i++ {
		ret[i] = element
	}
	return ret
}
