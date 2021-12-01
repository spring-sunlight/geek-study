package example_hash_set_map

import "container/list"

type LRUCache struct {
	capacity  int
	innerMap  map[int]*list.Element
	innerList *list.List
}

type entry struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		innerMap:  make(map[int]*list.Element),
		innerList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if l, ok := this.innerMap[key]; !ok {
		return -1
	} else {
		this.innerList.MoveToBack(l)
		return l.Value.(*entry).val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if l, ok := this.innerMap[key]; !ok {
		e := &entry{key: key, val: value}
		this.innerList.PushBack(e)
		this.innerMap[key] = this.innerList.Back()
		if len(this.innerMap) > this.capacity {
			delete(this.innerMap, this.innerList.Front().Value.(*entry).key)
			this.innerList.Remove(this.innerList.Front())
		}
	} else {
		l.Value.(*entry).val = value
		this.innerList.MoveToBack(l)
	}
}
