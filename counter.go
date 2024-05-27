package q

import "fmt"

// Counter is a generic counter data structure that stores unique elements of type T and counts occurences.
type Counter[T comparable] struct {
	data map[T]int
}

// NewCounter creates a new Counter and returns a pointer to it.
// Time complexity: O(1).
func NewCounter[T comparable](elements ...T) *Counter[T] {
	counter := &Counter[T]{data: make(map[T]int)}
	counter.Add(elements...)
	return counter
}

// Add adds one or more elements to the counter.
// Time complexity: O(n), where n is the number of elements being added.
func (c *Counter[T]) Add(elements ...T) {
	for _, element := range elements {
		c.data[element]++
	}
}

// Remove removes an element from the counter.
// Time complexity: O(1).
func (c *Counter[T]) Remove(element T) {
	delete(c.data, element)
}

// Contains checks if an element is present in the counter.
// Time complexity: O(1).
func (c *Counter[T]) Contains(element T) bool {
	_, exists := c.data[element]
	return exists
}

func (c *Counter[T]) Count(element T) int {
	if c.Contains(element) {
		return c.data[element]
	}
	return 0
}

// Len returns the number of elements in the counter.
// Time complexity: O(1).
func (c *Counter[T]) Len() int {
	return len(c.data)
}

// Clear removes all elements from the counter.
// Time complexity: O(1).
func (c *Counter[T]) Clear() {
	c.data = make(map[T]int)
}

// Elements returns a slice containing all the elements in the counter.
// Time complexity: O(n), where n is the number of elements in the counter.
func (c *Counter[T]) Elements() []T {
	elements := make([]T, 0, len(c.data))
	for element := range c.data {
		elements = append(elements, element)
	}
	return elements
}

// Equal checks if the counter is equal to another counter.
// Time complexity: O(n), where n is the number of elements in the counter.
func (c *Counter[T]) Equal(other *Counter[T]) bool {
	if c.Len() != other.Len() {
		return false
	}
	for element, count := range c.data {
		if other.Count(element) != count {
			return false
		}
	}
	return true
}

func (c *Counter[T]) String() string {
	return fmt.Sprintf("%v", c.data)
}
