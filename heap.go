package q

// Heap is a generic implementation of a heap data structure.
type Heap[M any] struct {
	less func(a, b M) bool
	data []M
}

// NewHeap creates a new instance of Heap with the specified less function.
// Time complexity: O(1).
func NewHeap[M any](less func(a, b M) bool) *Heap[M] {
	return &Heap[M]{less: less}
}

// Push adds one or more values to the heap.
// Time complexity: O(log n) for each value, where n is the number of elements in the heap.
func (h *Heap[M]) Push(values ...M) {
	for _, value := range values {
		h.data = append(h.data, value)
		h.up(len(h.data) - 1)
	}
}

// Pop removes and returns the top element from the heap.
// Time complexity: O(log n), where n is the number of elements in the heap.
func (h *Heap[M]) Pop() M {
	if len(h.data) == 0 {
		var m M
		return m
	}
	h.swap(0, len(h.data)-1)
	value := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return value
}

// Len returns the number of elements in the heap.
// Time complexity: O(1).
func (h *Heap[M]) Len() int {
	return len(h.data)
}

// Top returns the top element of the heap without removing it.
// Time complexity: O(1).
func (h *Heap[M]) Top() M {
	if len(h.data) == 0 {
		var m M
		return m
	}
	return h.data[0]
}

// Empty returns true if the heap is empty, false otherwise.
// Time complexity: O(1).
func (h *Heap[M]) Empty() bool {
	return len(h.data) == 0
}

// up moves the element at index i up the heap until the heap property is satisfied.
// Time complexity: O(log n), where n is the number of elements in the heap.
func (h *Heap[M]) up(i int) {
	for {
		j := (i - 1) / 2
		if i == j || !h.less(h.data[i], h.data[j]) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// down moves the element at index i down the heap until the heap property is satisfied.
// Time complexity: O(log n), where n is the number of elements in the heap.
func (h *Heap[M]) down(i int) {
	for {
		j := 2*i + 1
		if j >= len(h.data) {
			break
		}
		if j+1 < len(h.data) && h.less(h.data[j+1], h.data[j]) {
			j++
		}
		if h.less(h.data[i], h.data[j]) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

// swap swaps the elements at indices i and j in the heap.
// Time complexity: O(1).
func (h *Heap[M]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
