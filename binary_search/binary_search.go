package main

func BinarySearch(vals []int, target int) int {

	left := 0
	right := len(vals) - 1
	middle := (right - left) / 2

	if target > vals[right] || target < vals[left] {
		return -1
	}

	for vals[middle] != target {
		if vals[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
		middle = (right + left) / 2
	}
	return middle
}