package list

import "sort"

type ByteSlice []byte

func (x ByteSlice) Len() int           { return len(x) }
func (x ByteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x ByteSlice) Less(i, j int) bool { return x[i] < x[j] }

// InsertionSort implementation of in-place insertion sort (bubble sort) algorithm
// Complexity: O(n^2)
//
//	for j(index) ← 2 to n
//	  insert key A[j] into the (already sorted) sub-array A[1 .. j-1].
//	by pairwise key-swaps down to its right position
func InsertionSort(data sort.Interface, lo, hi int) {
	for i := lo + 1; i < hi; i++ {
		for j := i; j > lo && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func merge[K int | byte](left, right []K) []K {
	i, j, lenLeft, lenRight := 0, 0, len(left), len(right)
	var result []K
	for i < lenLeft && j < lenRight {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// MergeSort implementation of merge sort algorithm
// Complexity: O(n*ln(n))
//
//	require double size of memory
//	Algorithm:
//	 1. Split array on a half (merge_sort) recursively till 1 elements array
//	 2. merge two pre sorted sub array (merge) (on the deepest level array contents just one element)
func MergeSort[K int | byte](data []K) []K {
	n := len(data)
	mid := n >> 1
	if n > 1 {
		left := MergeSort(data[:mid])
		right := MergeSort(data[mid:])
		return merge(left, right)
	}

	return data
}

// shiftDown implements the heap property on data[lo:hi].
func shiftDown(data sort.Interface, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		// get the largest child
		if child+1 < hi && data.Less(child, child+1) {
			child++
		}
		if !data.Less(root, child) {
			return
		}
		data.Swap(root, child)
		root = child
	}
}

// HeapSort is implementation of the heap sort algorithm
// after n iterations the Heap is empty
//
//	every iteration involves a swap and a max_heapify
//	operation; hence it takes O(log n) time
//
// Complexity: O(n*ln(n))
func HeapSort(data sort.Interface) {
	lo, hi := 0, data.Len()
	// Build heap with the greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		shiftDown(data, i, hi)
	}
	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		data.Swap(lo, i)
		shiftDown(data, lo, i)
	}
}
