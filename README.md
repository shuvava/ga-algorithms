# Algorithms
LetCode algorithms implementation on Golang

## List Algorithms

package [list](./pkg/list/doc.go) contains algorithm using list as main data structure  

## Linked List

includes implementation of Linked List

```go
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shuvava/go-algorithms/pkg/linkedlist"
)

func main() {
	dll := linkedlist.NewDoubleLinkedList[int]()
	for _, v := range []int{1, 3, 5, 6} {
		dll.Add(v) // add a new val
	}
	iter := dll.Iterator() //get DoubleLinkedList iterator
	result := make([]string, 0)
	for iter.HasNext() {
		v := iter.Next()
		result = append(result, strconv.FormatInt(int64(v), 10))
	}
	fmt.Println(strings.Join(result[:], ", "))
}
// Output: 1, 3, 5, 6
```

## Collections

### Ordered Map

```go
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shuvava/go-algorithms/pkg/collections"
)

func main() {
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
}
// Output: 1, 3, 5, 6
```