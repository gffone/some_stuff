package main

type Queue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func NewLL() Queue {
	dummy := &Node{}
	return Queue{dummy, dummy, 0}
}

func (q *Queue) Push(val int) {
	
	newNode := Node{
		val: val,
	}
	
	q.size++
	q.tail.next = &newNode
	q.tail = q.tail.next
}

func (l *Queue) Pop() int {
	
	if q.head == q.tail {
		return -1
	}

	q.size--
	q.head = q.head.next
	return q.head.val
}

func (q *Queue) Reverse() {
	
	if q.size < 2 {
		return
	}

	var prev *Node

	q.tail = q.head.next
	prev = q.tail
	q.head = prev.next
	prev.next = nil
	for {
		if q.head.next == nil {
			dummy := &Node{}
			q.head.next, prev, q.head = prev, q.head, dummy
			q.head.next = prev
			return
		}
		q.head.next, prev, q.head = prev, q.head, q.head.next
	}
}
