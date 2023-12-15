package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDistinct(t *testing.T) {
	source := []int{1, 2, 3, 1, 2}
	target := Distinct(source)
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})
	assert.Equal(t, target, []int{1, 2, 3})
}
