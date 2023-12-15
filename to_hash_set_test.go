package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestToHashSet(t *testing.T) {
	source := []int{1, 2, 3, 3}
	set := ToHashSet(source)
	assert.Equal(t, len(set), 3)
	_, ok := set[1]
	assert.Equal(t, ok, true)
	_, ok = set[2]
	assert.Equal(t, ok, true)
	_, ok = set[3]
	assert.Equal(t, ok, true)
}
