package linq

func Count[T any](source []T, f func(x T) bool) int {
	cnt := 0
	for _, s := range source {
		if f(s) {
			cnt++
		}
	}
	return cnt
}
