package permutations_test

import (
	"fmt"

	"github.com/shuvava/go-algorithms/pkg/permutations"
)

func ExampleGetAllPermutations() {
	a := []int{1, 2, 3}
	iter := permutations.GetAllPermutations(a)
	var res [][]int
	for iter.HasNext() {
		item := iter.Next()
		res = append(res, item)
	}
	fmt.Println(res)
	// Output: [[1 2 3] [2 1 3] [3 1 2] [1 3 2] [2 3 1] [3 2 1]]
}
