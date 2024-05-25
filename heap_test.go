package q

import (
	"fmt"
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
	assert.Equal("[3 5 7]", h.String())
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

// ExampleHeap_Push demonstrates how to push elements onto the heap.
func ExampleHeap_Push() {
	h := NewHeap(func(a, b int) bool {
		return a < b
	})
	h.Push(5, 2, 7, 1, 9)
	fmt.Println(h.data)
	// Output: [1 2 7 5 9]
}

// ExampleHeap_Pop demonstrates how to pop the top element from the heap.
func ExampleHeap_Pop() {
	h := NewHeap(func(a, b int) bool {
		return a < b
	})
	h.Push(5, 2, 7, 1, 9)
	top := h.Pop()
	fmt.Println(top)
	fmt.Println(h.data)
	// Output:
	// 1
	// [2 5 7 9]
}

// ExampleHeap_Top demonstrates how to retrieve the top element without removing it.
func ExampleHeap_Top() {
	h := NewHeap(func(a, b string) bool {
		return a < b
	})
	h.Push("banana", "apple", "orange", "grape")
	top := h.Top()
	fmt.Println(top)
	fmt.Println(h.data)
	// Output:
	// apple
	// [apple banana orange grape]
}

// ExampleHeap_Empty demonstrates how to check if the heap is empty.
func ExampleHeap_Empty() {
	h := NewHeap(func(a, b int) bool {
		return a < b
	})
	fmt.Println(h.Empty())
	h.Push(1, 2, 3)
	fmt.Println(h.Empty())
	// Output:
	// true
	// false
}

// ExampleHeap_Len demonstrates how to get the number of elements in the heap.
func ExampleHeap_Len() {
	h := NewHeap(func(a, b int) bool {
		return a < b
	})
	h.Push(5, 2, 7, 1, 9)
	fmt.Println(h.Len())
	// Output: 5
}
