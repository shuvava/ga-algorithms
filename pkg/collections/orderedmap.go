package collections

import (
	"github.com/shuvava/go-algorithms/pkg/linkedlist"
	"github.com/shuvava/go-algorithms/pkg/list"
)

// Pair is tuple of key and value of the OrderedMap
type Pair[K, V any] struct {
	Key   K
	Value V
}

type iteratorOrderedMap[K comparable, V any] struct {
	iter list.Iterator[Pair[K, V]]
}

type OrderedMap[K comparable, V any] struct {
	items map[K]*linkedlist.Item[Pair[K, V]]
	link  *linkedlist.DoubleLinkedList[Pair[K, V]]
}

// NewOrderedMap creates a new OrderedMap
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		items: make(map[K]*linkedlist.Item[Pair[K, V]]),
		link:  linkedlist.NewDoubleLinkedList[Pair[K, V]](),
	}
}

// Get lookup or the given key, and returns the value associated with it if it is found
func (c *OrderedMap[K, V]) Get(key K) (value V, found bool) {
	if item, present := c.items[key]; present {
		return item.Value.Value, true
	}
	return
}

// Set adds or replaces key in OrderedMap if key was present in the map func returns previous value
func (c *OrderedMap[K, V]) Set(key K, value V) (val V, replaced bool) {
	if item, present := c.items[key]; present {
		oldValue := item.Value.Value
		item.Value = Pair[K, V]{Key: key, Value: value}
		return oldValue, true
	}

	item := c.link.Add(Pair[K, V]{Key: key, Value: value})
	c.items[key] = item

	return
}

// Delete try to remove key from OrderedMap and return value of the key if it was present
func (c *OrderedMap[K, V]) Delete(key K) (val V, replaced bool) {
	if item, present := c.items[key]; present {
		oldValue := item.Value.Value
		c.link.Del(item)
		return oldValue, true
	}
	return
}

// Len return number of elements in the OrderedMap
func (c *OrderedMap[K, V]) Len() int {
	return len(c.items)
}

// Iterator return iterator interface for the OrderedMap
func (c *OrderedMap[K, V]) Iterator() list.Iterator[Pair[K, V]] {
	return &iteratorOrderedMap[K, V]{
		iter: c.link.Iterator(),
	}
}

// Next returns the next Item value of the OrderedMap
func (i *iteratorOrderedMap[K, V]) Next() Pair[K, V] {
	return i.iter.Next()
}

// HasNext returns true if there are value to be read
func (i *iteratorOrderedMap[K, V]) HasNext() bool {
	return i.iter.HasNext()
}
