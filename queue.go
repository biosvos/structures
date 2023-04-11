package structures

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any](elements ...T) *Queue[T] {
	return &Queue[T]{
		elements: elements,
	}
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() T {
	ret := q.elements[0]
	q.elements = q.elements[1:]
	return ret
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
