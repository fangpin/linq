package linq

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestChunk(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	chunks := Chunk(source, 2)
	assert.Equal(t, len(chunks), 3)
	assert.Equal(t, chunks[0], []int{1, 2})
	assert.Equal(t, chunks[1], []int{3, 4})
	assert.Equal(t, chunks[2], []int{5})

	source = []int{1, 2, 3, 4}
	chunks = Chunk(source, 2)
	assert.Equal(t, len(chunks), 2)
	assert.Equal(t, chunks[0], []int{1, 2})
	assert.Equal(t, chunks[1], []int{3, 4})
}
