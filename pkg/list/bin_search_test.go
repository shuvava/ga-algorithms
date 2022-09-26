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
		{"even size middle target Not Fond", []int{1, 3, 5, 6}, 2, 1},
		{"even size most right target Not Fond", []int{1, 3, 5, 6}, 9, 4},
		{"even size most left target Not Fond", []int{1, 3, 5, 6}, 0, 0},
		{"even size most right target", []int{1, 3, 5, 9}, 9, 3},
		{"even size most left target", []int{1, 3, 5, 9}, 1, 0},
		{"even size middle target #1", []int{1, 3, 5, 9}, 5, 2},
		{"even size middle target #2", []int{1, 3, 5, 9}, 3, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, found := list.BinSearch(tc.list, tc.target)
			if r != tc.expected {
				t.Errorf("got %d, expected %d", r, tc.expected)
			}
			if found && tc.list[r] != tc.target {
				t.Errorf("for target:%d looking up in  array[%v] expected %d got %d", tc.target, tc.list, tc.expected, r)
			}
		})
	}
}

func TestBinSelectLeft(t *testing.T) {
	testCases := []TestCase{
		{"even size middle target Not Fond", []int{1, 3, 5, 6}, 2, 1},
		{"even size most right target Not Fond", []int{1, 3, 5, 6}, 9, 4},
		{"even size most left target Not Fond", []int{1, 3, 5, 6}, 0, 0},
		{"even size most right target", []int{1, 3, 5, 9}, 9, 3},
		{"even size most left target", []int{1, 3, 5, 9}, 1, 0},
		{"even size middle target #1", []int{1, 3, 5, 9}, 5, 2},
		{"even size middle target #2", []int{1, 3, 5, 9}, 3, 1},
		{"should return left most index", []int{0, 1, 1, 1, 2}, 1, 1},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := list.BinSelectLeft(tc.list, tc.target)
			if r != tc.expected {
				t.Errorf("got %d, expected %d", r, tc.expected)
			}
		})
	}
}

func FuzzBinSelectLeft(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, target byte) {
		aa := make([]int, len(a))
		for i, v := range a {
			aa[i] = int(v)
		}
		sort.Ints(aa)
		for i, v := range aa {
			a[i] = byte(v)
		}
		expected := sort.SearchInts(aa, int(target))
		r := list.BinSelectLeft(a, target)
		if r != expected {
			t.Errorf("for target:%d looking up in  array[%v] expected %d got %d", target, a, expected, r)
		}
	})
}

func TestBinSelectRight(t *testing.T) {
	testCases := []TestCase{
		{"even size middle target Not Fond", []int{1, 3, 5, 6}, 2, 1},
		{"even size most right target Not Fond", []int{1, 3, 5, 6}, 9, 4},
		{"even size most left target Not Fond", []int{1, 3, 5, 6}, 0, 0},
		{"even size most right target", []int{1, 3, 5, 9}, 9, 4},
		{"even size most left target", []int{1, 3, 5, 9}, 1, 1},
		{"even size middle target #1", []int{1, 3, 5, 9}, 5, 3},
		{"even size middle target #2", []int{1, 3, 5, 9}, 3, 2},
		{"should return left most index", []int{0, 1, 1, 1, 2}, 1, 4},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := list.BinSelectRight(tc.list, tc.target)
			if r != tc.expected {
				t.Errorf("got %d, expected %d", r, tc.expected)
			}
		})
	}
}
