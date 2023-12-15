package linq

import (
	"sort"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestValues(t *testing.T) {
	m := map[int]string{
		1: "1",
		2: "2",
	}
	values := Values(m)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	assert.Equal(t, values, []string{"1", "2"})
}
