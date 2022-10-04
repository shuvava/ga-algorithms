package linkedlist

import "github.com/shuvava/go-algorithms/pkg/list"

// Item is linked list item data structure
type Item[K any] struct {
	Prev *Item[K]
	Next *Item[K]
	Key  K
}

// DoubleLinkedList single thread implementation of Double Linked List
type DoubleLinkedList[K any] struct {
	root *Item[K]
}

type iterator[K any] struct {
	dll     *DoubleLinkedList[K]
	current *Item[K]
	list.Iterator[K]
}

// NewItem creates new Item of LinkedList object
func NewItem[K any](key K) *Item[K] {
	llist := Item[K]{
		Key: key,
	}
	llist.Prev = nil
	llist.Next = nil
	return &llist
}

// NewDoubleLinkedList creates double linked list object
func NewDoubleLinkedList[K any]() *DoubleLinkedList[K] {
	root := Item[K]{}
	root.Prev = &root
	root.Next = &root
	return &DoubleLinkedList[K]{
		root: &root,
	}
}

// Add adds a new Item to the end of the DoubleLintedList and returns reference on it
// Complexity: O(1)
func (dll *DoubleLinkedList[K]) Add(key K) *Item[K] {
	item := NewItem(key)
	last := dll.root.Prev
	item.Prev, item.Next = last, dll.root
	last.Next = item
	dll.root.Prev = item
	return item
}

// Del removes Item from the DoubleLinkedList
// Complexity: O(1)
func (dll *DoubleLinkedList[K]) Del(item *Item[K]) {
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
func (dll *DoubleLinkedList[K]) Iterator() list.Iterator[K] {
	return &iterator[K]{
		current: dll.root,
		dll:     dll,
	}
}

// Len returns number of items in the DoubleLinkedList
func (dll *DoubleLinkedList[K]) Len() int {
	root, item := dll.root, dll.root
	n := 0
	for item.Next != root {
		n++
		item = item.Next
	}
	return n
}

// Next returns the next Item value of the DoubleLinkedList
func (i *iterator[K]) Next() K {
	i.current = i.current.Next
	return i.current.Key
}

// HasNext returns true if there are value to be read
func (i *iterator[K]) HasNext() bool {
	return i.current.Next != i.dll.root
}
