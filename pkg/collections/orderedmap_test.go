package collections_test

import (
	"strconv"
	"testing"

	"github.com/shuvava/go-algorithms/pkg/collections"
)

func TestOrderedMap(t *testing.T) {
	t.Run("New OrderedMap should return 0 Len", func(t *testing.T) {
		om := collections.NewOrderedMap[string, int]()
		l := om.Len()
		if l != 0 {
			t.Errorf("Length of empty OrderedMap should be 0")
		}
	})
	t.Run("add should increase size of OrderedMap", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		om := collections.NewOrderedMap[string, int]()
		for k := range a {
			om.Set(strconv.FormatInt(int64(k), 10), k)
		}
		if len(a) != om.Len() {
			t.Errorf("Length of DoubleLinkedList is not equal to test array: len(arr)=%d len(om)=%d", len(a), om.Len())
		}
	})
	t.Run("add should maintain order", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		om := collections.NewOrderedMap[string, int]()
		for _, k := range a {
			om.Set(strconv.FormatInt(int64(k), 10), k)
		}
		iter := om.Iterator()
		for _, k := range a {
			item := iter.Next()
			if item.Key != strconv.FormatInt(int64(k), 10) || item.Value != k {
				t.Errorf("error in maintaineng map order")
			}
		}
	})
}
