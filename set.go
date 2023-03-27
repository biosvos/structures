package structures

import "sort"

type Set[T Identifier] struct {
	elements map[uint64]T
}

func NewSet[T Identifier]() *Set[T] {
	return &Set[T]{
		elements: map[uint64]T{},
	}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.elements[element.Id()]
	return ok
}

func (s *Set[T]) Add(element T) {
	s.elements[element.Id()] = element
}

func (s *Set[T]) Delete(element T) {
	delete(s.elements, element.Id())
}

func (s *Set[T]) Slice() []T {
	var keys []uint64
	for key := range s.elements {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var ret []T
	for _, key := range keys {
		ret = append(ret, s.elements[key])
	}
	return ret
}
