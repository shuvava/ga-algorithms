package permutations

import "github.com/shuvava/go-algorithms/pkg/list"

// permutations is a state for heap algorithm
type permutations[T any] struct {
	items []T
	len   int
	stack []int
	pos   int
	next  []T
}

// GetAllPermutations returns all permutations of provided elements
// Heap's algorithm
//
//	https://en.wikipedia.org/wiki/Heap%27s_algorithm
//	https://www.baeldung.com/java-array-permutations
func GetAllPermutations[T any](elements []T) list.Iterator[[]T] {
	l := len(elements)
	state := &permutations[T]{
		items: elements,
		stack: make([]int, l),
		len:   l,
		pos:   0,
	}
	state.next = make([]T, state.len)
	copy(state.next, elements)
	return state
}

func (p *permutations[T]) nextPermutations() (res []T) {
	for p.pos < p.len {
		if p.stack[p.pos] < p.pos {
			if p.pos%2 == 0 {
				p.items[p.pos], p.items[0] = p.items[0], p.items[p.pos]
			} else {
				p.items[p.pos], p.items[p.stack[p.pos]] = p.items[p.stack[p.pos]], p.items[p.pos]
			}
			p.stack[p.pos] += 1
			// Simulate recursive call reaching the base case
			// by bringing the pointer to the base case analog in the array
			p.pos = 0
			res = make([]T, p.len)
			copy(res, p.items)
			return res
		} else {
			// Reset the state and simulate popping
			//  the stack by incrementing the pointer.
			p.stack[p.pos] = 0
			p.pos += 1
		}
	}
	return
}

// Next returns the next permutation
func (p *permutations[T]) Next() []T {
	next := p.next
	p.next = p.nextPermutations()
	return next
}

// HasNext returns true if there are value to be read
func (p *permutations[T]) HasNext() bool {
	return p.next != nil
}
