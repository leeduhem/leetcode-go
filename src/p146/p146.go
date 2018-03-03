package p146

import (
	"container/list"
)

type Item struct {
	key, value int
}

type LRUCache struct {
	capacity int
	cache    map[int](*list.Element)
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	this := LRUCache{
		capacity: capacity,
		cache: map[int](*list.Element) {},
		list: list.New(),
	}

	return this
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.cache[key]; ok {
		this.list.MoveToFront(e)
		return e.Value.(*Item).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	// Update existed element
	if e, ok := this.cache[key]; ok {
		this.list.MoveToFront(e)
		e.Value = &Item{key: key, value: value}
		return
	}

	// Add new element
	if this.capacity > 0 {
		this.capacity--
	} else {
		e := this.list.Back()
		delete(this.cache, e.Value.(*Item).key)
		this.list.Remove(e)
	}

	this.list.PushFront(&Item{key: key, value: value})
	this.cache[key] = this.list.Front()
}
