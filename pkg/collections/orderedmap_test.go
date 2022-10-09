package collections_test

import (
	"fmt"
	"strconv"
	"strings"
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
	t.Run("PopLast should return the oldest element in map", func(t *testing.T) {
		a := []int{1, 3, 5, 6}
		om := collections.NewOrderedMap[string, int]()
		for _, k := range a {
			om.Set(strconv.FormatInt(int64(k), 10), k)
		}
		for _, k := range a {
			el, found := om.PopLast()
			if !found {
				t.Errorf("should found a element")
			}
			if el != k {
				t.Errorf("PopLast shoul return latest element in array")
			}
		}
	})
	t.Run("PopLast should not found for empty map", func(t *testing.T) {
		om := collections.NewOrderedMap[string, int]()
		_, found := om.PopLast()
		if found {
			t.Errorf("PopLast should return false for empty map")
		}
	})
}

func ExampleNewOrderedMap() {
	om := collections.NewOrderedMap[string, int]()
	for _, v := range []int{1, 3, 5, 6} {
		om.Set(strconv.FormatInt(int64(v), 10), v) // add new element to the OrderedMap
	}
	iter := om.Iterator() //get OrderedMap iterator
	result := make([]string, 0)
	for iter.HasNext() {
		kv := iter.Next()
		result = append(result, kv.Key)
	}
	fmt.Println(strings.Join(result[:], ", "))
	// Output: 1, 3, 5, 6
}
