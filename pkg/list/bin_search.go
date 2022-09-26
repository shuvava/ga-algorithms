package list

// BinSearch implementation of binary search on golang
// this func return -1 with error if target is not found
func BinSearch[K int | int64 | uint | byte](a []K, target K) (int, bool) {
	l := 0
	h := len(a)
	for h > l {
		m := (h + l) >> 1
		if a[m] < target {
			l = m + 1
		} else {
			h = m
		}
	}
	return l, l < len(a) && a[l] == target
}

// BinSelectLeft returns the index where to insert item target in list a,
// assuming a is sorted.
//
//	The return value i is such that all e in a[:i] have e < target, and all e in
//
// a[i:] have e >= target.
//
//	So if target already appears in the list, a.insert(i, target) will
//
// insert just before the leftmost x already there.
func BinSelectLeft[K int | int64 | uint | byte](a []K, target K) int {
	l := 0
	h := len(a)
	for l < h {
		m := (h + l) >> 1
		if a[m] < target {
			l = m + 1
		} else {
			h = m
		}
	}
	return l
}

// BinSelectRight returns the index where to insert item target in list a, assuming a is sorted.
//
//	The return value i is such that all e in a[:i] have e <= target, and all e in
//	a[i:] have e > target.  So if x already appears in the list, a.insert(i, target) will
//	insert just after the rightmost x already there.
func BinSelectRight[K int | int64 | uint | byte](a []K, target K) int {
	l := 0
	h := len(a)
	for l < h {
		m := (h + l) >> 1
		if target < a[m] {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}
