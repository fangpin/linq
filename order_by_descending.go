package linq

import "golang.org/x/exp/constraints"

func OrderByDescending[T constraints.Ordered](source []T, comparer func(l, r T) bool) []T {
	return Reverse(OrderBy(source, comparer))
}
