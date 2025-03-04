package merge_sort_list

type ListNode struct {
	val  int
	next *ListNode
}

func GetMid(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func ListNodeMerge(a *ListNode, b *ListNode) *ListNode {
	zeroNode := &ListNode{-1, nil}
	curr := zeroNode
	for a != nil && b != nil {
		if a.val <= b.val {
			curr.next = a
			a = a.next
		} else {
			curr.next = b
			b = b.next
		}
		curr = curr.next
	}
	if a == nil {
		curr.next = b
	}
	if b == nil {
		curr.next = a
	}
	return zeroNode.next
}

func ListNodeMergeSort(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}
	mid := GetMid(head)
	b := mid.next
	mid.next = nil
	return ListNodeMerge(ListNodeMergeSort(head), ListNodeMergeSort(b))
}
