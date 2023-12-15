package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestDefaultIfEmpty(t *testing.T) {
	test := []int{1}
	assert.Equal(t, DefaultIfEmpty(test), []int{1})
	test = nil
	assert.Equal(t, DefaultIfEmpty(test), []int{0})
	test = []int{}
	assert.Equal(t, DefaultIfEmpty(test), []int{0})
}
