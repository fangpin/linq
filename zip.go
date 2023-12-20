package linq

type Pair[T, R any] struct {
	First  T
	Second R
}

func Zip[T, R any](left []T, right []R) []Pair[T, R] {
	if len(left) != len(right) {
		panic(InvalidParam)
	}
	result := []Pair[T, R]{}
	for i := 0; i < len(left); i++ {
		result = append(result, Pair[T, R]{left[i], right[i]})
	}
	return result
}
