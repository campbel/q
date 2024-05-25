package q

import "fmt"

// List represents a generic linked list.
type List[M any] struct {
	length int64
	head   *Node[M]
	tail   *Node[M]
}

// Node represents a node in the linked list.
type Node[M any] struct {
	value M
	next  *Node[M]
	prev  *Node[M]
}

// NewList creates a new List and initializes it with the given elements.
// Time complexity: O(n), where n is the number of elements.
func NewList[M any](elements ...M) *List[M] {
	l := &List[M]{}
	for _, e := range elements {
		l.PushRight(e)
	}
	return l
}

// Push adds values to the list.
// Time complexity: O(n), where n is the number of values.
func (l *List[M]) Push(values ...M) {
	l.PushRight(values...)
}

// Pop removes and returns the last value from the list.
// Time complexity: O(1).
func (l *List[M]) Pop() M {
	return l.PopRight()
}

// PushRight adds values to the end of the list.
// Time complexity: O(n), where n is the number of values.
func (l *List[M]) PushRight(values ...M) {
	for _, v := range values {
		l.length++
		node := &Node[M]{value: v}
		if l.head == nil {
			l.head = node
			l.tail = node
		} else {
			l.tail.next = node
			node.prev = l.tail
			l.tail = node
		}
	}
}

// PopRight removes and returns the last value from the list.
// Time complexity: O(1).
func (l *List[M]) PopRight() M {
	l.length--
	if l.tail == nil {
		var m M
		return m
	}
	node := l.tail
	l.tail = node.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	return node.value
}

// Extend appends other lists to the current list.
// Time complexity: O(m), where m is the total number of elements in all the lists.
func (l *List[M]) Extend(lists ...*List[M]) {
	for _, other := range lists {
		if other.length == 0 {
			return
		}
		if l.head == nil {
			l.head = other.head
			l.tail = other.tail
		} else {
			l.tail.next = other.head
			other.head.prev = l.tail
			l.tail = other.tail
		}
		l.length += other.length
	}
}

// PopLeft removes and returns the first value from the list.
// Time complexity: O(1).
func (l *List[M]) PopLeft() M {
	l.length--
	if l.head == nil {
		var m M
		return m
	}
	node := l.head
	l.head = node.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	return node.value
}

// PushLeft adds values to the beginning of the list.
// Time complexity: O(n), where n is the number of values.
func (l *List[M]) PushLeft(values ...M) {
	for _, v := range values {
		l.length++
		node := &Node[M]{value: v}
		if l.head == nil {
			l.head = node
			l.tail = node
		} else {
			l.head.prev = node
			node.next = l.head
			l.head = node
		}
	}
}

// Reverse reverses the order of the list.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Reverse() {
	for n := l.head; n != nil; n = n.prev {
		n.next, n.prev = n.prev, n.next
	}
	l.head, l.tail = l.tail, l.head
}

// String returns a string representation of the list.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) String() string {
	return fmt.Sprintf("%v", l.Elements())
}

// Each applies a callback function to each value in the list.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Each(callback func(M)) {
	for n := l.head; n != nil; n = n.next {
		callback(n.value)
	}
}

// Find returns the first value in the list that satisfies the callback function.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Find(callback func(M) bool) M {
	for n := l.head; n != nil; n = n.next {
		if callback(n.value) {
			return n.value
		}
	}
	var m M
	return m
}

// Elements returns a slice containing all the values in the list.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Elements() []M {
	result := make([]M, 0, l.length)
	for n := l.head; n != nil; n = n.next {
		result = append(result, n.value)
	}
	return result
}

// Filter returns a new list containing only the values that satisfy the callback function.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Filter(callback func(M) bool) *List[M] {
	result := &List[M]{}
	for n := l.head; n != nil; n = n.next {
		if callback(n.value) {
			result.PushRight(n.value)
		}
	}
	return result
}

// All returns true if all values in the list satisfy the callback function, false otherwise.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) All(callback func(M) bool) bool {
	for n := l.head; n != nil; n = n.next {
		if !callback(n.value) {
			return false
		}
	}
	return true
}

// Any returns true if at least one value in the list satisfies the callback function, false otherwise.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Any(callback func(M) bool) bool {
	for n := l.head; n != nil; n = n.next {
		if callback(n.value) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the value in the list, or -1 if not found.
// Time complexity: O(n), where n is the number of elements in the list.
func IndexOf[M comparable](list *List[M], value M) int {
	index := 0
	for n := list.head; n != nil; n = n.next {
		if n.value == value {
			return index
		}
		index++
	}
	return -1
}

// Remove removes all occurrences of the value from the list.
// Time complexity: O(n), where n is the number of elements in the list.
func Remove[M comparable](list *List[M], value M) {
	for n := list.head; n != nil; n = n.next {
		if n.value == value {
			if n == list.head {
				list.head = n.next
				if n.next != nil {
					n.next.prev = nil
				}
			} else if n == list.tail {
				list.tail = n.prev
				n.prev.next = nil
			} else {
				n.prev.next = n.next
				n.next.prev = n.prev
			}
			list.length--
		}
	}
}

// Sort returns a new sorted list using the provided less function.
// Time complexity: O(n log n), where n is the number of elements in the list.
func (l *List[M]) Sort(less func(M, M) bool) *List[M] {
	if l.length < 2 {
		return l
	}
	left, pivot, right := l.partition(less)
	return stitch(left.Sort(less), pivot, right.Sort(less))
}

// partition partitions the list into three parts: left, pivot, and right.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) partition(less func(M, M) bool) (*List[M], M, *List[M]) {
	pivot := l.head.value
	left := NewList[M]()
	right := NewList[M]()
	for n := l.head.next; n != nil; n = n.next {
		if less(n.value, pivot) {
			left.PushRight(n.value)
		} else {
			right.PushRight(n.value)
		}
	}
	return left, pivot, right
}

// stitch combines the left, pivot, and right lists into a single list.
// Time complexity: O(n), where n is the total number of elements in the left, pivot, and right lists.
func stitch[M any](left *List[M], pivot M, right *List[M]) *List[M] {
	if left.length == 0 {
		return Join(NewList[M](pivot), right)
	}
	left.PushRight(pivot)
	return Join(left, right)
}

// IsSorted returns true if the list is sorted in non-decreasing order according to the provided less function, false otherwise.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) IsSorted(less func(M, M) bool) bool {
	for n := l.head; n != nil && n.next != nil; n = n.next {
		if less(n.next.value, n.value) {
			return false
		}
	}
	return true
}

// Copy returns a new list with the same values as the original list.
// Time complexity: O(n), where n is the number of elements in the list.
func (l *List[M]) Copy() *List[M] {
	result := &List[M]{}
	for n := l.head; n != nil; n = n.next {
		result.PushRight(n.value)
	}
	return result
}

// Len returns the length of the list.
// Time complexity: O(1).
func (l *List[M]) Len() int {
	return int(l.length)
}

// Len64 returns the length of the list as an int64.
// Time complexity: O(1).
func (l *List[M]) Len64() int64 {
	return l.length
}

// Map applies a callback function to each value in the list and returns a new list with the results.
// Time complexity: O(n), where n is the number of elements in the list.
func Map[M any, N any](l *List[M], callback func(M) N) *List[N] {
	result := &List[N]{}
	for n := l.head; n != nil; n = n.next {
		result.PushRight(callback(n.value))
	}
	return result
}

// Reduce applies a callback function to each value in the list and returns a single accumulated value.
// Time complexity: O(n), where n is the number of elements in the list.
func Reduce[M any, N any](l *List[M], callback func(N, M) N, initial N) N {
	acc := initial
	for n := l.head; n != nil; n = n.next {
		acc = callback(acc, n.value)
	}
	return acc
}

// Equal returns true if the two lists are equal, false otherwise.
// Time complexity: O(n), where n is the number of elements in the list.
func Equal[M comparable](a, b *List[M]) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.length != b.length {
		return false
	}
	for na, nb := a.head, b.head; na != nil && nb != nil; na, nb = na.next, nb.next {
		if na.value != nb.value {
			return false
		}
	}
	return true
}

// Join combines multiple lists into a single list.
// Time complexity: O(m), where m is the total number of elements in all the lists.
func Join[M any](lists ...*List[M]) *List[M] {
	if len(lists) == 0 {
		return NewList[M]()
	}

	if len(lists) == 1 {
		return lists[0]
	}

	l := lists[0]
	for _, list := range lists[1:] {
		if list.length == 0 {
			continue
		}
		l.tail.next = list.head
		list.head.prev = l.tail
		l.tail = list.tail
		l.length += list.length
	}

	return l
}
