package q

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		s.data[element] = struct{}{}
	}
}

func (s *Set[T]) Remove(element T) {
	delete(s.data, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.data[element]
	return exists
}

func (s *Set[T]) Len() int {
	return len(s.data)
}

func (s *Set[T]) Clear() {
	s.data = make(map[T]struct{})
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.data))
	for element := range s.data {
		elements = append(elements, element)
	}
	return elements
}

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

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for element := range s.data {
		if !other.Contains(element) {
			result.Add(element)
		}
	}
	return result
}
