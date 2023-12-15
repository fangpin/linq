package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAll(t *testing.T) {
	source := []int{2, 4, 5}
	assert.Equal(t, All(source, func(x int) bool {
		return x%2 == 0
	}), false)
	source = []int{2, 4, 6}
	assert.Equal(t, All(source, func(x int) bool {
		return x%2 == 0
	}), true)
}
