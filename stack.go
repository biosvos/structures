package structures

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return true
}

func (s *Stack[T]) Pop() T {
	lastIndex := len(s.elements) - 1
	ret := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return ret
}

func (s *Stack[T]) Push(str T) {
	s.elements = append(s.elements, str)
}

func (s *Stack[T]) Peek() T {
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
