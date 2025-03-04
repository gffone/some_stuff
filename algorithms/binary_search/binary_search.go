package main

func BinarySearch(vals []int, target int) int {

	left := 0
	right := len(vals) - 1

	if target > vals[right] || target < vals[left] {
		return -1
	}

	for left <= right {
		middle := (right + left) / 2

		if vals[middle] == target {
			return middle
		}

		if vals[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
