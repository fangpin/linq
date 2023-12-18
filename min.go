package linq

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](source []T) T {
	if len(source) == 0 {
		panic("empty source")
	}
	min := First(source, func(x T) bool { return true })
	for i := 1; i < len(source); i++ {
		if source[i] < min {
			min = source[i]
		}
	}
	return min
}