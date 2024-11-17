package lru

import (
	"container/list"
	"iter"
)

type LRUcache[K comparable, V any] interface {
	Put(key K, value V)
	Get(key K) V
	Size() int

	All() iter.Seq2[K, V]
}

type Node[K comparable, V any] struct {
	key   K
	value V
}

type cacheImpl[K comparable, V any] struct {
	linkedList   *list.List
	keyToElement map[K]*list.Element
	capacity     int
	defaultValue V
}

func (c *cacheImpl[K, V]) GetNodeFromElement(element *list.Element) *Node[K, V] {
	switch v := element.Value.(type) {
	case *Node[K, V]:
		return v
	default:
		panic("")
	}
}

func (c *cacheImpl[K, V]) extractLatest() {
	del := c.linkedList.Back()

	delete(c.keyToElement, c.GetNodeFromElement(del).key)

	c.linkedList.Remove(del)

}

func (c *cacheImpl[K, V]) Put(key K, value V) {
	if link, ok := c.keyToElement[key]; ok {
		node := c.GetNodeFromElement(link)
		node.value = value
		c.linkedList.MoveToFront(link)
		return
	}

	if c.Size() == c.capacity {
		c.extractLatest()
	} else {
		node := &Node[K, V]{key, value}
		c.keyToElement[key] = c.linkedList.PushFront(node)

	}
}

func (c *cacheImpl[K, V]) Get(key K) V {
	if link, ok := c.keyToElement[key]; !ok {
		return c.defaultValue
	} else {
		return c.GetNodeFromElement(link).value
	}
}

func (c *cacheImpl[K, V]) Size() int {
	return c.linkedList.Len()
}

func (c *cacheImpl[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		cur := c.linkedList.Front()

		for range c.Size() {
			n := c.GetNodeFromElement(cur)

			if !yield(n.key, n.value) {
				return
			}

			cur = cur.Next()
		}
	}
}

func NewCache[K comparable, V any](capacity int, defaultValue V) *cacheImpl[K, V] {
	return &cacheImpl[K, V]{
		list.New(),
		make(map[K]*list.Element, capacity),
		capacity,
		defaultValue,
	}
}

/*
func foo() {
	l:= NewCache(1, 0)

	it:=l.All()

	for k, v := range it{
		if k==v{
			break
		}
	}

	сахар на ->

	it(func(i1, i2 int)bool{
		return i1==i2
	})

}
*/
