package linq

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func OrderBy[T constraints.Ordered](source []T, comparer func(l, r T) bool) []T {
	sort.Slice(source, func(i, j int) bool {
		return comparer(source[i], source[j])
	})
	return source
}
