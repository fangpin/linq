package linq

import "golang.org/x/exp/constraints"

func OrderDescending[T constraints.Ordered](source []T) []T {
	return Reverse(Order(source))
}
