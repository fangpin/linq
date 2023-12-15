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
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})
	assert.Equal(t, target, []int{1, 2})
}
