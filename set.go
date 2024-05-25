package q

// Set is a generic set data structure that stores unique elements of type T.
type Set[T comparable] struct {
	data map[T]struct{}
}

// NewSet creates a new Set and returns a pointer to it.
// Time complexity: O(1).
func NewSet[T comparable](elements ...T) *Set[T] {
	set := &Set[T]{data: make(map[T]struct{})}
	set.Add(elements...)
	return set
}

// Add adds one or more elements to the set.
// Time complexity: O(n), where n is the number of elements being added.
func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.data[element] = struct{}{}
	}
}

// Remove removes an element from the set.
// Time complexity: O(1).
func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

// Contains checks if an element is present in the set.
// Time complexity: O(1).
func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

// Len returns the number of elements in the set.
// Time complexity: O(1).
func (s *Set[T]) Len() int {
	return len(s.data)
}

// Clear removes all elements from the set.
// Time complexity: O(1).
func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

// Elements returns a slice containing all the elements in the set.
// Time complexity: O(n), where n is the number of elements in the set.
func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

// Union returns a new set that is the union of the current set and another set.
// Time complexity: O(n), where n is the total number of elements in both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		result.Add(element)
	}
	for element := range other.data {
		result.Add(element)
	}
	return result
}

// Intersection returns a new set that is the intersection of the current set and another set.
// Time complexity: O(n), where n is the number of elements in the smaller set.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

// Difference returns a new set that contains the elements present in the current set but not in another set.
// Time complexity: O(n), where n is the number of elements in the current set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
