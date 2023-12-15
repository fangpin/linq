package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestExcept(t *testing.T) {
	l := []int{1, 2, 3, 4}
	r := []int{4, 5}
	target := Except(l, r)
	sort.Slice(target, func(i, j int) bool {
		return target[i] < target[j]
	})
	assert.Equal(t, target, []int{1, 2, 3})
}
