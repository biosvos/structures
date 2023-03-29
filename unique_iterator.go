package structures

type UniqueIterator[K comparable, T Identifier[K]] struct {
	checker  *Set[K, T]
	iterator *Iterator[T]
}

func NewUniqueIterator[K comparable, T Identifier[K]](elements ...T) *UniqueIterator[K, T] {
	ret := UniqueIterator[K, T]{
		checker:  NewSet[K, T](),
		iterator: NewIterator[T](),
	}
	ret.Add(elements...)
	return &ret
}

func (u *UniqueIterator[K, T]) HasNext() bool {
	return u.iterator.HasNext()
}

func (u *UniqueIterator[K, T]) Next() T {
	return u.iterator.Next()
}

func (u *UniqueIterator[K, T]) Add(elements ...T) {
	for _, element := range elements {
		if !u.checker.Has(element) {
			u.checker.Add(element)
			u.iterator.Add(element)
		}
	}
}
