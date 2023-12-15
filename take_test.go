package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTake(t *testing.T) {
	source := []int{1, 2, 3, 4}
	target := Take(source, 3)
	assert.Equal(t, target, []int{1, 2, 3})
}
