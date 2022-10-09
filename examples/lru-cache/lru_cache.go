package lru_cache

import (
	"github.com/shuvava/go-algorithms/pkg/collections"
	"github.com/shuvava/go-algorithms/pkg/list"
)

// LRUCache is design and implement a data structure for Least Recently Used (LRU) cache.
//
//	It should support the following operations: get and put.
//
//	 get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//	 put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity,
//	 it should invalidate the least recently used item before inserting a new item.
//
//	The cache is initialized with a positive capacity.
type LRUCache[K comparable, V any] struct {
	size int
	dict *collections.OrderedMap[K, V]
}

// NewLRUCache create new instance of LRU cache with capacity equal to size
func NewLRUCache[K comparable, V any](size int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		size: size,
		dict: collections.NewOrderedMap[K, V](),
	}
}

// Get return value of the cached key if found
func (c *LRUCache[K, V]) Get(key K) (value V, found bool) {
	return c.dict.Get(key)
}

// Delete removes element from cache
func (c *LRUCache[K, V]) Delete(key K) (val V, found bool) {
	return c.dict.Delete(key)
}

// Set add or update element in the cache
//
//	 update operation do not change order of elements
//		if cache capacity is reached the limit the oldest element will be removed from the cache
func (c *LRUCache[K, V]) Set(key K, value V) (val V, replaced bool) {
	val, replaced = c.dict.Set(key, value)
	if c.dict.Len() > c.size {
		_, _ = c.dict.PopLast()
	}
	return val, replaced
}

// Len return number of elements in the LRUCache
func (c *LRUCache[K, V]) Len() int {
	return c.dict.Len()
}

// Iterator return iterator interface for the LRUCache
func (c *LRUCache[K, V]) Iterator() list.Iterator[collections.Pair[K, V]] {
	return c.dict.Iterator()
}
