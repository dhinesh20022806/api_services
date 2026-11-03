// Package set is data struture allow avoid duplicates
package set

// Set struct
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet is function used to create new Set
func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		elements: make(map[T]struct{}),
	}

	s.Add(items...)

	return s
}

// Add method used to append new items
func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.elements[item] = struct{}{}
	}
}

// Remove method used to remove by the item
func (s *Set[T]) Remove(item T) {
	delete(s.elements, item)
}

// Has method check wheather items exists or not
func (s *Set[T]) Has(item T) bool {
	_, exists := s.elements[item]
	return exists
}

// Size method give length of the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Clear method reassign elements
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// ToSlice convert custom set to slice go native data structure
func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))

	for item := range s.elements {
		slice = append(slice, item)
	}
	return slice
}
