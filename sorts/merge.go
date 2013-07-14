package sorts

func MergeSort(toSort []int) []int {
	length := len(toSort)

	if length <= 1 {
		return toSort
	}

	return Merge(MergeSort(toSort[:length/2]), MergeSort(toSort[(length/2)+1:]))
}

func Merge(a []int, b []int) []int {
	merged := make([]int, 1)

	for len(a) > 0 && len(b) > 0 {
		if a[0] < b[0] {
			merged = append(merged, a[0])
			Dequeue(&a)
		} else {
			merged = append(merged, b[0])
			Dequeue(&b)
		}
	}

	var leftOver []int

	if len(a) > 0 {
		leftOver = a
	} else {
		leftOver = b
	}

	for _, num := range leftOver {
		merged = append(merged, num)
	}

	return merged
}

func Dequeue(a *[]int) {
	*a = (*a)[1:]
}
