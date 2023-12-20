package linq

import "golang.org/x/exp/constraints"

func MaxBy[T any, K constraints.Ordered](source []T, com func(x T) K) T {
	if len(source) == 0 {
		panic(EmptySource)
	}
	maxId, maxKey := 0, com(source[0])
	for i := 1; i < len(source); i++ {
		key := com(source[i])
		if key > maxKey {
			maxKey = key
			maxId = i
		}
	}
	return source[maxId]
}
