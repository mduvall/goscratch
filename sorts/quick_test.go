package sorts

import (
	"sort"
	"testing"
)

func Test_QuickSortSorts(t *testing.T) {
	toSort := []int{3, 5, 9, 11, 1, 4, 5}
	sorted := QuickSort(toSort)

	if !sort.IntsAreSorted(sorted) {
		t.Errorf("Should be sorted man")
	}
}
