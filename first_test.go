package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFirst(t *testing.T) {
	source := []int{1, 2, 3, 4}
	assert.Equal(t, First(source, func(x int) bool {
		return x%2 == 0
	}), 2)
	assert.Panic(t, func() {
		First(source, func(x int) bool {
			return x > 5
		})
	}, "not found satisfied item")
}
