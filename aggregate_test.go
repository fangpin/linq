package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAggregagte(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6}
	sum := Aggregate(source, 0, func(prev, now int) int {
		return prev + now
	})
	assert.Equal(t, sum, 21)
}
