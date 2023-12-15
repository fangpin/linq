package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestKeys(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
	}
	keys := Keys(m)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	assert.Equal(t, keys, []int{1, 2})
}
