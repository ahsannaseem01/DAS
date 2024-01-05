package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	i := []int{1, 2, 3, 100, 1009, 2000}
	tar := 1009

	act := BinarySearch(i, tar)
	if !act {
		t.Errorf("target %v not found", tar)
	}
}
