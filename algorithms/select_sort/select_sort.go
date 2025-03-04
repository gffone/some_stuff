package select_sort

func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i += 1 {
		minIndex := i
		for j := i + 1; j < length; j += 1 {
			if arr[j] < arr[minIndex] {
				t := arr[minIndex]
				arr[minIndex] = arr[j]
				arr[j] = t
				minIndex = j
			}
		}
	}
	return arr
}
