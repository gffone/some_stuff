package merge_sort

func Merge(arr []int, lo int, mid int, hi int) {
	leftSize := mid - lo + 1
	rightSize := hi - mid
	left := make([]int, leftSize)
	right := make([]int, rightSize)
	for i := 0; i < leftSize; i++ {
		left[i] = arr[lo+i]
	}
	for j := 0; j < rightSize; j++ {
		right[j] = arr[mid+j+1]
	}
	i, j, k := 0, 0, lo
	for k <= hi && i < leftSize && j < rightSize {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < leftSize && k <= hi {
		arr[k] = left[i]
		i++
		k++
	}
	for j < rightSize && k <= hi {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int, lo int, hi int) {
	if lo < hi {
		mid := (lo + hi) / 2
		MergeSort(arr, lo, mid)
		MergeSort(arr, mid+1, hi)
		Merge(arr, lo, mid, hi)
	}
}
