package list

// Iterator is an interface for generic Iterator
type Iterator[K any] interface {
	Next() K
	HasNext() bool
}
