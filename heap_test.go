package q

import (
	"math/rand"
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

func TestHeapWithRandom(t *testing.T) {
	assert := assert.New(t)

	// Test NewHeap
	h := NewHeap(func(a, b int) bool {
		return a < b
	})

	// Test Push and Len
	count := 10000
	for i := 0; i < count; i++ {
		h.Push(rand.Int())
	}
	assert.Equal(count, h.Len())
	assert.False(h.Empty())

	// Test Pop
	previous := h.Top()
	for i := 0; i < count; i++ {
		assert.True(previous <= h.Top(), "previous: %d, top: %d", previous, h.Top())
		previous = h.Pop()
	}
	assert.Zero(h.Pop())
	assert.Zero(h.Top())
	assert.True(h.Empty())
}
