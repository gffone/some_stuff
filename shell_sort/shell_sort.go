package shell_sort

func ShellSort(arr []int) []int {
	n := len(arr)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		h /= 3
	}
	return arr
}
