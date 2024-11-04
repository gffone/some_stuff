package main

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func NewLL() LinkedList {
	dummy := &Node{}
	return LinkedList{dummy, dummy, 0}
}

func (l *LinkedList) Push(val int) {
	newNode := Node{
		val: val,
	}
	l.size++
	l.tail.next = &newNode
	l.tail = l.tail.next
}

func (l *LinkedList) Pop() int {
	if l.head == l.tail {
		return -1
	}
	l.size--
	l.head = l.head.next
	return l.head.val
}

func (l *LinkedList) Reverse() {

	if l.size < 2 {
		return
	}

	var prev *Node

	l.tail = l.head.next
	prev = l.tail
	l.head = prev.next
	prev.next = nil
	for {
		if l.head.next == nil {
			dummy := &Node{}
			l.head.next, prev, l.head = prev, l.head, dummy
			l.head.next = prev
			return
		}
		l.head.next, prev, l.head = prev, l.head, l.head.next
	}
}

func main() {
}
