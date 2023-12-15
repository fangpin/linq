package linq

func DefaultIfEmpty[T any](source []T) []T {
	if len(source) == 0 {
		source = []T{zero[T]{}.inner}
	}
	return source
}
