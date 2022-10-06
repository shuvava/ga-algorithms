package linkedlist

import "github.com/shuvava/go-algorithms/pkg/list"

// Item is linked list item data structure
type Item[V any] struct {
	Prev  *Item[V]
	Next  *Item[V]
	Value V
}

// DoubleLinkedList single thread implementation of Double Linked List
type DoubleLinkedList[V any] struct {
	root *Item[V]
}

type iterator[V any] struct {
	dll     *DoubleLinkedList[V]
	current *Item[V]
	list.Iterator[V]
}

// NewItem creates new Item of LinkedList object
func NewItem[V any](val V) *Item[V] {
	llist := Item[V]{
		Value: val,
	}
	llist.Prev = nil
	llist.Next = nil
	return &llist
}

// NewDoubleLinkedList creates double linked list object
func NewDoubleLinkedList[V any]() *DoubleLinkedList[V] {
	root := Item[V]{}
	root.Prev = &root
	root.Next = &root
	return &DoubleLinkedList[V]{
		root: &root,
	}
}

// Add adds a new Item to the end of the DoubleLintedList and returns reference on it
// Complexity: O(1)
func (dll *DoubleLinkedList[V]) Add(val V) *Item[V] {
	item := NewItem(val)
	last := dll.root.Prev
	item.Prev, item.Next = last, dll.root
	last.Next = item
	dll.root.Prev = item
	return item
}

// Del removes Item from the DoubleLinkedList
// Complexity: O(1)
func (dll *DoubleLinkedList[V]) Del(item *Item[V]) {
	if item == nil {
		return
	}
	prev, next := item.Prev, item.Next
	prev.Next = next
	next.Prev = prev
	item.Next = nil
	item.Prev = nil
}

// Iterator return iterator interface for the DoubleLinkedList
func (dll *DoubleLinkedList[V]) Iterator() list.Iterator[V] {
	return &iterator[V]{
		current: dll.root,
		dll:     dll,
	}
}

// Len returns number of items in the DoubleLinkedList
func (dll *DoubleLinkedList[V]) Len() int {
	root, item := dll.root, dll.root
	n := 0
	for item.Next != root {
		n++
		item = item.Next
	}
	return n
}

// Next returns the next Item value of the DoubleLinkedList
func (i *iterator[V]) Next() V {
	i.current = i.current.Next
	return i.current.Value
}

// HasNext returns true if there are value to be read
func (i *iterator[V]) HasNext() bool {
	return i.current.Next != i.dll.root
}
