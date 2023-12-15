package linq

import (
	"strconv"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSelect(t *testing.T) {
	source := []int{1, 2, 3}
	target := Select(source, func(x int) string { return strconv.Itoa(x * x) })
	assert.Equal(t, target, []string{"1", "4", "9"})
}
