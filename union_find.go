package structures

type UnionFind[T comparable] struct {
	leaders map[T]T
}

func NewUnionFind[T comparable]() *UnionFind[T] {
	return &UnionFind[T]{
		leaders: map[T]T{},
	}
}

func (f *UnionFind[T]) Root(follower T) T {
	_, ok := f.leaders[follower]
	if !ok {
		return follower
	}
	f.leaders[follower] = f.Root(f.leaders[follower])
	return f.leaders[follower]
}

func (f *UnionFind[T]) Merge(leader T, follower T) {
	leader = f.Root(leader)
	follower = f.Root(follower)
	if leader == follower {
		return
	}
	f.leaders[follower] = leader
}

func (f *UnionFind[T]) IsSame(a T, b T) bool {
	a = f.Root(a)
	b = f.Root(b)
	return a == b
}
