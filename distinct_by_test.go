package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDistinckBy(t *testing.T) {
	source := []int{1, 2, 3, 4}
	target := DistinctBy(source, func(x int) int {
		return x % 2
	})
	sort.Slice(source, func(i, j int) bool {
		return source[i] < source[j]
	})
	assert.Equal(t, target, []int{1, 2})
}
