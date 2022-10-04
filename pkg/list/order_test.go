package list_test

import (
	"testing"

	"github.com/shuvava/go-algorithms/pkg/list"
)

func TestReverse(t *testing.T) {
	t.Run("test of revers for even number of elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		expected := []int{4, 3, 2, 1}
		list.Reverse(a)
		for i, v := range a {
			if v != expected[i] {
				t.Errorf("Reverse func return incorrect result for index %d", i)
			}
		}
	})
	t.Run("test of revers for uneven number of elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		list.Reverse(a)
		for i, v := range a {
			if v != expected[i] {
				t.Errorf("Reverse func return incorrect result for index %d", i)
			}
		}
	})
}

func TestPopLast(t *testing.T) {
	t.Run("PopLast should remove last element", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		l := len(a)
		got := make([]int, l)
		for i := 0; i < l; i++ {
			v, a1 := list.PopLast(a)
			a = a1
			got[i] = v
		}
		for i := 0; i < l; i++ {
			if expected[i] != got[i] {
				t.Errorf("PopLast func return incorrect result for index %d", i)
			}
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete should remove reference types", func(t *testing.T) {
		a := [][]int{{1}, {2}, {3}, {4}, {5}}
		l := len(a)
		got := list.Delete(a, 2)
		if len(got) != l-1 {
			t.Errorf("delete function return incorrect number of elemnts")
		}
		if got[1][0] != 2 {
			t.Errorf("delete function broke the first part of slice")
		}
		if got[2][0] != 4 {
			t.Errorf("delete function broke the second part of slice")
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Delete should remove reference types", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5}
		l := len(a)
		el, got := list.Pop(a, 2)
		if len(got) != l-1 {
			t.Errorf("Pop function return incorrect number of elemnts")
		}
		if el != 3 {
			t.Errorf("Pop return incorrect item")
		}
		if got[1] != 2 {
			t.Errorf("Pop function broke the first part of slice")
		}
		if got[2] != 4 {
			t.Errorf("delete function broke the second part of slice")
		}
	})
}
