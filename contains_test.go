package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestContains(t *testing.T) {
	source := []int{1, 2, 3}
	assert.Equal(t, Contains(source, 1), true)
	assert.Equal(t, Contains(source, 4), false)
}
