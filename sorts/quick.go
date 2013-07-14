package sorts

import (
	"math/rand"
	"time"
)

func QuickSort(toSort []int) []int {
	length := len(toSort)

	// Base case to catch the recursion at
	if length <= 1 {
		return toSort
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pivotIdx := r.Intn(length)
	pivot := toSort[pivotIdx]

	// Rip the pivot out
	toSort = append(toSort[:pivotIdx], toSort[pivotIdx+1:]...)
	left, right := []int{}, []int{}

	for _, num := range toSort {
		if num >= pivot {
			right = append(right, num)
		} else {
			left = append(left, num)
		}
	}

	return append(QuickSort(left), append([]int{pivot}, QuickSort(right)...)...)
}
