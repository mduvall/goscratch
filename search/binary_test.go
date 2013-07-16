package search

import (
	"testing"
)

func Test_BinarySearchWorks(t *testing.T) {
	a := []int{1, 4, 8, 18, 28, 39, 54, 299, 1888}
	if BinarySearch(a, len(a)-1, 0, 4) == -1 {
		t.Errorf("Binary search failed")
	}

	if BinarySearch(a, len(a)-1, 0, 5) != -1 {
		t.Errorf("Binary search failed")
	}
}
