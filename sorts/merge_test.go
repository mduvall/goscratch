package sorts

import (
	"sort"
	"testing"
)

func Test_MergeSortSortsCorrectly(t *testing.T) {
	toSort := []int{4, 5, 5, 3, 2, 1, 2, 4, 5, 6, 7, 4}
	sorted := MergeSort(toSort)

	if !sort.IntsAreSorted(sorted) {
		t.Errorf("The ints aren't sorted bro")
	}
}
