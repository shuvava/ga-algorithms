package list_test

import (
	"sort"
	"testing"

	"github.com/shuvava/go-algorithms/pkg/list"
)

type TestCase struct {
	name     string
	list     []int
	target   int
	expected int
}

func TestBinarySearch(t *testing.T) {
	testCases := []TestCase{
		{"even size middle target Not Fond", []int{1, 3, 5, 6}, 2, -1},
		{"even size most right target Not Fond", []int{1, 3, 5, 6}, 9, -1},
		{"even size most left target Not Fond", []int{1, 3, 5, 6}, 0, -1},
		{"even size most right target", []int{1, 3, 5, 9}, 9, 3},
		{"even size most left target", []int{1, 3, 5, 9}, 1, 0},
		{"even size middle target #1", []int{1, 3, 5, 9}, 5, 2},
		{"even size middle target #2", []int{1, 3, 5, 9}, 3, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, err := list.BinSearch(tc.list, tc.target)
			if tc.expected >= 0 && r != tc.expected {
				t.Errorf("got %d, expected %d", r, tc.expected)
			}
			if tc.expected == -1 && err == nil {
				t.Errorf("get %d, expected NOT nil error", r)
			} else if tc.expected >= 0 && err != nil {
				t.Errorf("get %s, expected nil", err)
			}
		})
	}
}

func FuzzBinSearch(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, target byte) {
		aa := make([]int, len(a))
		for i, v := range a {
			aa[i] = int(v)
		}
		sort.Ints(aa)
		expected := sort.SearchInts(aa, int(target))
		r, err := list.BinSearch(a, target)
		if a[expected] == target && r != expected {
			t.Errorf("got %d, expected %d", r, expected)
		}
		if a[expected] != target && err == nil {
			t.Errorf("get %d, expected NOT nil error", r)
		} else if a[expected] == target && err != nil {
			t.Errorf("get %s, expected nil", err)
		}
	})
}
