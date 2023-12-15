package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDistinct(t *testing.T) {
	source := []int{1, 2, 3, 1, 2}
	target := Distinct(source)
	assert.Equal(t, target, []int{1, 2, 3})
}
