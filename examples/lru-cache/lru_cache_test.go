package lru_cache_test

import (
	"fmt"
	lru_cache "github.com/shuvava/go-algorithms/examples/lru-cache"
	"strconv"
	"strings"
)

func ExampleNewLRUCache() {
	a := []int{1, 2, 3, 3, 2, 4}
	size := 2
	cache := lru_cache.NewLRUCache[string, int](size)
	for _, v := range a {
		cache.Set(strconv.FormatInt(int64(v), 10), v)
	}
	result := make([]string, 0)
	iter := cache.Iterator()
	for iter.HasNext() {
		e := iter.Next()
		result = append(result, e.Key)
	}
	fmt.Println(strings.Join(result[:], ", "))
	// Output: 3, 4
}
