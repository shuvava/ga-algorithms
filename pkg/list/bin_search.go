package list

import "fmt"

// BinSearch implementation of binary search on golang
// this func return -1 with error if target is not found
func BinSearch[K int | int64 | uint | byte](a []K, target K) (int, error) {
	l := 0
	h := len(a)
	for h > l {
		m := (h + l) >> 1
		if a[m] == target {
			return m, nil
		}
		if a[m] > target {
			h = m
		} else {
			l = m + 1
		}
	}
	return -1, fmt.Errorf("target value (%d) not found", target)
}
