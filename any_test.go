package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAny(t *testing.T) {
	source := []int{2, 4, 6}
	assert.Equal(t, Any(source, func(x int) bool { return x%2 == 1 }), false)
	source = []int{2, 4, 5}
	assert.Equal(t, Any(source, func(x int) bool { return x%2 == 1 }), true)
}
