package linq

import (
	"sort"
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIntersectBy(t *testing.T) {
	source := []int{1, 2, 3, 4}
	keys := []string{"1", "3", "5"}
	target := IntersectBy(source, keys, func(x int) string {
		return strconv.Itoa(x)
	})
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})
	assert.Equal(t, target, []int{1, 3})
}
