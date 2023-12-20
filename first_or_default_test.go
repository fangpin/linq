package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFirstOrDefault(t *testing.T) {
	source := []int{1, 2, 3, 4}
	assert.Equal(t, FirstOrDefault(source, func(x int) bool {
		return x > 5
	}), 0)
	assert.Equal(
		t,
		FirstOrDefault([]int{3, 5, 7, 8}, func(x int) bool {
			return x > 5
		}), 7)
}
