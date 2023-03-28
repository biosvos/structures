package structures

type Iterator[T any] struct {
	elements []T
}

func NewIterator[T any](elements ...T) *Iterator[T] {
	return &Iterator[T]{elements: elements}
}

func (i *Iterator[T]) HasNext() bool {
	return len(i.elements) > 0
}

func (i *Iterator[T]) Next() T {
	ret := i.elements[0]
	i.elements = i.elements[1:]
	return ret
}

func (i *Iterator[T]) Add(elements ...T) {
	i.elements = append(i.elements, elements...)
}
