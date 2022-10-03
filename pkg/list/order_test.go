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
