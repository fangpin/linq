package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCount(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	cnt := Count(source, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, cnt, 4)
}
