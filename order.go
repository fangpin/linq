package linq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Order[T constraints.Ordered](source []T) []T {
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})
	return source
}
