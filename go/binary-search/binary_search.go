package binarysearch

func SearchInts(list []int, key int) int {
	left := 0
	right := len(list) - 1
	for left <= right {
		m := (right + left) / 2
		if list[m] > key {
			right = m - 1
		} else if list[m] < key {
			left = m + 1
		} else {
			return m
		}
	}

	return -1
}
