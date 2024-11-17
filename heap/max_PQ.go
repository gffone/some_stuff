package heap

import "fmt"

type MaxPQ struct {
	heap []int
	n    int
}

func (pq *MaxPQ) Insert(node int) {
	pq.heap[pq.n] = node
	fmt.Println(pq.heap, pq.n)
	pq.n += 1
	if pq.n > 1 {
		Swap(pq.heap, 0, pq.n-1)

		BuildHeap(pq.heap, pq.n)
	}
}

func (pq *MaxPQ) DelMax() {
	if pq.n == 0 {
		return
	}
	Swap(pq.heap, 0, pq.n-1)

	pq.heap[pq.n-1] = 0
	pq.n -= 1
	fmt.Println("del:", pq.heap, pq.n)
	BuildHeap(pq.heap, pq.n)
}

func (pq *MaxPQ) Heap() []int {
	return pq.heap
}

func (pq *MaxPQ) Size() int {
	return pq.n
}

func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func Heapify(tree []int, n int, i int) {
	if i >= n {
		return
	}
	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}
	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}
	if max != i {
		Swap(tree, max, i)
		Heapify(tree, n, max)
	}
}

func BuildHeap(tree []int, n int) {
	lastNode := n - 1
	var parent int = (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		Heapify(tree, n, i)
	}
}

func HeapSort(tree []int, n int) {
	BuildHeap(tree, n)
	for i := n - 1; i > 0; i-- {
		Swap(tree, i, 0)
		Heapify(tree, i, 0)
	}
}
