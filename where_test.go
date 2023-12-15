package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestWhere(t *testing.T) {
	source := []int{1, 2, 3, 4}
	target := Where(source, func(x int) bool { return x%2 == 1 })
	assert.Equal(t, target, []int{1, 3})
}
