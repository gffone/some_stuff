package insert_sort

func InsertSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i += 1 {
		for j := i; j > 0 && arr[j] < arr[j-1]; j -= 1 {
			t := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = t
		}
	}
	return arr
}
