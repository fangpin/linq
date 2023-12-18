package linq

import "golang.org/x/exp/constraints"

func MinBy[T any, K constraints.Ordered](source []T, com func(x T) K) T {
	if len(source) == 0 {
		panic("empty source")
	}
	minId, minKey := 0, com(source[0])
	for i := 1; i < len(source); i++ {
		key := com(source[i])
		if key < minKey {
			minKey = key
			minId = i
		}
	}
	return source[minId]
}
