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
