package q

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	assert := assert.New(t)

	// Test NewHeap
	h := NewHeap(func(a, b int) bool {
		return a < b
	})

	// Test Push and Len
	h.Push(5, 3, 7)
	assert.Equal(3, h.Len())
	assert.Equal(3, h.Top())
	assert.False(h.Empty())

	// Test Pop
	assert.Equal(3, h.Pop())
	assert.Equal(5, h.Pop())
	assert.Equal(7, h.Pop())
	assert.Zero(h.Pop())
	assert.True(h.Empty())
}
