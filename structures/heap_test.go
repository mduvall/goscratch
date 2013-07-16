package structures

import (
	"testing"
)

func Test_BuildHeapWorks(t *testing.T) {
	vals := []int{5, 4, 2, 1, 43, 48, 2, 5, 55, 2, 18}
	h := BuildHeap(vals)

	for len(h.Elements) > 1 {
		h.DeleteMin()
	}
}
