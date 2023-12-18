package linq

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](source []T) T {
	if len(source) == 0 {
		panic("empty source")
	}
	max := First(source, func(x T) bool { return true })
	for i := 1; i < len(source); i++ {
		if source[i] > max {
			max = source[i]
		}
	}
	return max
}
