package structures

type Identifier[T comparable] interface {
	Identify() T
}
