package search

func BinarySearch(i []int, high int, low int, key int) int {
	if low > high {
		return -1
	}

	middle := (high + low) / 2

	if i[middle] == key {
		return i[middle]
	}

	if i[middle] > key {
		return BinarySearch(i, middle-1, low, key)
	} else {
		return BinarySearch(i, high, middle+1, key)
	}
}
