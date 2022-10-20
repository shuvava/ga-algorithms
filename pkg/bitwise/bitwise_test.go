package bitwise_test

import (
	"fmt"
	"testing"

	"github.com/shuvava/go-algorithms/pkg/bitwise"
)

func TestGetMask(t *testing.T) {
	testCases := []struct {
		bitNum   byte
		expected int
	}{
		{16, 65_535},
		{8, 255},
		{1, 1},
		{0, 0},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("GetMask should return %d for %d bite", test.expected, test.bitNum), func(t *testing.T) {
			got := bitwise.GetMask[int](test.bitNum)
			if got != test.expected {
				t.Errorf("expected %d got %d", test.expected, got)
			}
		})
	}
}

func TestGetBit(t *testing.T) {
	testCases := []struct {
		value    int
		bitNum   byte
		expected bool
	}{
		{7, 0, true},
		{7, 1, true},
		{7, 2, true},
		{7, 3, false},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("GetBit should return %t for value %d and bit %d", test.expected, test.value, test.bitNum), func(t *testing.T) {
			got := bitwise.GetBit[int](test.value, test.bitNum)
			if got != test.expected {
				t.Errorf("expected %t got %t", test.expected, got)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	testCases := []struct {
		value    int
		bitNum   byte
		expected int
	}{
		{0, 0, 1},
		{1, 1, 3},
		{3, 2, 7},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("GetBit should return %d for value %d and bit %d", test.expected, test.value, test.bitNum), func(t *testing.T) {
			got := bitwise.SetBit[int](test.value, test.bitNum)
			if got != test.expected {
				t.Errorf("expected %d got %d", test.expected, got)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	testCases := []struct {
		value    int
		bitNum   byte
		expected int
	}{
		{1, 0, 0},
		{3, 1, 1},
		{7, 2, 3},
		{7, 3, 7},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("GetBit should return %d for value %d and bit %d", test.expected, test.value, test.bitNum), func(t *testing.T) {
			got := bitwise.ClearBit[int](test.value, test.bitNum)
			if got != test.expected {
				t.Errorf("expected %d got %d", test.expected, got)
			}
		})
	}
}
