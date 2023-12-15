package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAverage(t *testing.T) {
	source := []int{1, 3, 5}
	assert.Equal(t, Average(source), 3.0)
}
