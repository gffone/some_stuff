package main

import (
	"sync"
)

type Queue struct {
	c    *sync.Cond
	head *Node
	tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

func NewQueue() *Queue {
	cond := sync.NewCond(&sync.Mutex{})
	dummy := &Node{}
	return &Queue{head: dummy, tail: dummy, c: cond}
}

func (q *Queue) Push(item int) {
	q.c.L.Lock()
	defer q.c.L.Unlock()
	q.size++

	newNode := Node{}
	q.tail.val = item
	q.tail.next = &newNode
	q.tail = q.tail.next
	q.c.Signal()
}

// Если очередь пуста, блокируется до момента,
// пока в очереди не появится элемент.
func (q *Queue) Pop() int {
	q.c.L.Lock()
	defer q.c.L.Unlock()
	for q.size == 0 {
		q.c.Wait()
	}
	q.size--
	q.head = q.head.next
	return q.head.val
}

func (q *Queue) Len() int {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	s := q.size

	return s
}
