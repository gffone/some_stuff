
package main

import (
	"fmt"
	"sync"
)


type Queue struct {
	c *sync.Cond
	size int
	firstNode *Node
	lastNode *Node
}

type Node struct{
	val int
	next *Node
}


func NewQueue() *Queue {
	cond:=sync.NewCond(&sync.Mutex{})
	var f  Node
	return &Queue{firstNode: &f, lastNode: &f, c:cond}
}


func (q *Queue) Put(item int) {
	q.c.L.Lock()
	defer q.c.L.Unlock()
	q.size++
	
	newNode:= Node{}
	q.lastNode.val = item
	q.lastNode.next = &newNode
	q.lastNode = q.lastNode.next
	q.c.Signal()
}


// Если очередь пуста, блокируется до момента,
// пока в очереди не появится элемент.
func (q *Queue) Get() int {
	q.c.L.Lock()
	defer q.c.L.Unlock()
	for q.size == 0{
		q.c.Wait()
	}
	q.size--
	v:=q.firstNode.val
	q.firstNode = q.firstNode.next
	return v
}

// Len возвращает количество элементов в очереди.
func (q *Queue) Len() int {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	s:=q.size

	return s
}

//конец решения

func main() {
	var wg sync.WaitGroup
	q := NewQueue()

	wg.Add(1)
	go func() {
		for i := range 100 {
			q.Put(i)
		}
		wg.Done()
	}()
	wg.Wait()

	total := 0

	wg.Add(1)
	go func() {
		for range 100 {
			total += q.Get()
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Put x100, Get x100, Total:", total)
}


