package list_test

import (
	"sort"
	"testing"

	"github.com/shuvava/go-algorithms/pkg/list"
)

func FuzzInsertionSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, target byte) {
		aa := make([]int, len(a))
		for i, v := range a {
			aa[i] = int(v)
		}
		sort.Ints(aa)
		list.InsertionSort(list.ByteSlice(a), 0, len(a))
		for i, v := range a {
			if aa[i] != int(v) {
				t.Errorf("for index %d value not match", i)
			}
		}
	})
}

func FuzzMergeSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, target byte) {
		aa := make([]int, len(a))
		for i, v := range a {
			aa[i] = int(v)
		}
		sort.Ints(aa)
		result := list.MergeSort(a)
		for i, v := range result {
			if aa[i] != int(v) {
				t.Errorf("for index %d value not match", i)
			}
		}
	})
}

func FuzzHeapSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, target byte) {
		aa := make([]int, len(a))
		for i, v := range a {
			aa[i] = int(v)
		}
		sort.Ints(aa)
		list.HeapSort(list.ByteSlice(a))
		for i, v := range a {
			if aa[i] != int(v) {
				t.Errorf("for index %d value not match", i)
			}
		}
	})
}
