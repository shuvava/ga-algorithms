package list

// Swap swaps two elements of provided array, without boundaries checks
func Swap[K any](x []K, i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Reverse in place reverse of the items in slice
// Complexity: O(N)
func Reverse[K any](data []K) {
	l := len(data)
	m := l >> 1
	for i := 0; i < m; i++ {
		Swap(data, i, l-i-1)
	}
}

// PopLast pop the latest element from slice
// Taken from: https://github.com/golang/go/wiki/SliceTricks
// Complexity: O(1)
func PopLast[K any](data []K) (K, []K) {
	return data[len(data)-1], data[:len(data)-1]
}

// Delete removes i elements from slice
// Taken from: https://github.com/golang/go/wiki/SliceTricks
// Complexity: O(N)
func Delete[K any](data []K, i int) []K {
	copy(data[i:], data[i+1:])

	return data[:len(data)-1]
}

// Pop pops from slice i element and return slice without this element
// Taken from: https://github.com/golang/go/wiki/SliceTricks
// Complexity: O(N)
func Pop[K any](data []K, i int) (K, []K) {
	l := len(data)
	if l-1 == i {
		return PopLast(data)
	}
	el := data[i]
	return el, Delete(data, i)
}
