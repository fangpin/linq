package linq

func TakeWhile[T any](source []T, takeCond func(x T) bool) []T {
	result := []T{}
	for _, s := range source {
		if takeCond(s) {
			result = append(result, s)
		} else {
			break
		}
	}
	return result
}
