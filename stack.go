package structures

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Pop() T {
	ret := s.lastElement()
	s.dropLastElement()
	return ret
}

func (s *Stack[T]) Push(element T) {
	s.append(element)
}

func (s *Stack[T]) Peek() T {
	return s.lastElement()
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) lastIndex() int {
	return s.Size() - 1
}

func (s *Stack[T]) dropLastElement() {
	s.elements = s.elements[:s.lastIndex()]
}

func (s *Stack[T]) append(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) lastElement() T {
	return s.elements[s.lastIndex()]
}
