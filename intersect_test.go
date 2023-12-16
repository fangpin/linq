package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIntersect(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	b := []int{2, 4, 5, 8, 9}
	c := Intersect(a, b)
	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})
	assert.Equal(t, c, []int{5, 9})
}
