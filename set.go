package structures

type Set[K comparable, T Identifier[K]] struct {
	elements map[K]T
}

func NewSet[K comparable, T Identifier[K]]() *Set[K, T] {
	return &Set[K, T]{
		elements: map[K]T{},
	}
}

func (s *Set[K, T]) Has(element T) bool {
	_, ok := s.elements[element.Identify()]
	return ok
}

func (s *Set[K, T]) Add(element T) {
	s.elements[element.Identify()] = element
}

func (s *Set[K, T]) Delete(element T) {
	delete(s.elements, element.Identify())
}

func (s *Set[K, T]) Slice() []T {
	var ret []T
	for _, value := range s.elements {
		ret = append(ret, value)
	}
	return ret
}
