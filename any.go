package linq

func Any[T any](source []T, predicator func(x T) bool) bool {
	for _, s := range source {
		if predicator(s) {
			return true
		}
	}
	return false
}
