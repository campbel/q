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
func (s *Counter[T]) Add(elements ...T) {
	for _, element := range elements {
		s.data[element]++
	}
}

// Remove removes an element from the counter.
// Time complexity: O(1).
func (s *Counter[T]) Remove(element T) {
	delete(s.data, element)
}

// Contains checks if an element is present in the counter.
// Time complexity: O(1).
func (s *Counter[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

func (s *Counter[T]) Count(element T) int {
	if s.Contains(element) {
		return s.data[element]
	}
	return 0
}

// Len returns the number of elements in the counter.
// Time complexity: O(1).
func (s *Counter[T]) Len() int {
	return len(s.data)
}

// Clear removes all elements from the counter.
// Time complexity: O(1).
func (s *Counter[T]) Clear() {
	s.data = make(map[T]int)
}

// Elements returns a slice containing all the elements in the counter.
// Time complexity: O(n), where n is the number of elements in the counter.
func (s *Counter[T]) Elements() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

// Equal checks if the counter is equal to another counter.
// Time complexity: O(n), where n is the number of elements in the counter.
func (s *Counter[T]) Equal(other *Counter[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for element, count := range s.data {
		if other.Count(element) != count {
			return false
		}
	}
	return true
}

func (s *Counter[T]) String() string {
	return fmt.Sprintf("%v", s.data)
}
