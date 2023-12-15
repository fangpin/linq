package linq

type _T[T any] struct {
	inner T
}

func DefaultIfEmpty[T any](source []T) []T {
	if len(source) == 0 {
		source = []T{_T[T]{}.inner}
	}
	return source
}
