package main

import (
	"fmt"
	"slices"
)

type Node struct {
	Key, Value int
}

type LRUCache struct {
	capacity int
	hash     map[int]int
	used     []Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     map[int]int{},
		used:     make([]Node, 0, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	// fmt.Println("Get(", key, ")")
	if value, ok := this.hash[key]; ok {
		idx := slices.IndexFunc(this.used, func(item Node) bool {
			return item.Key == key
		})
		this.used = slices.Delete(this.used, idx, idx+1)
		this.used = slices.Insert(this.used, 0, Node{
			Key:   key,
			Value: value,
		})
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// fmt.Println("Put(", key, ",", value, ")", this.used)
	if _, exists := this.hash[key]; exists {
		// fmt.Println(" [ ] Deleting", value, "from", this.used)
		idx := slices.IndexFunc(this.used, func(item Node) bool {
			return item.Key == key
		})
		this.used = slices.Delete(this.used, idx, idx+1)
		delete(this.hash, key)
		// fmt.Println(" [+] Delete", value, "from", this.used)
	}
	if len(this.used) == this.capacity {
		// fmt.Println(" [ ] Removing from", this.used)
		toDelete := this.used[this.capacity-1]
		this.used = slices.Delete(this.used, this.capacity-1, this.capacity)
		delete(this.hash, toDelete.Key) // Fuck
		// fmt.Println(" [+] Removed from", this.used)
	}
	this.used = slices.Insert(this.used, 0, Node{
		Key:   key,
		Value: value,
	})
	this.hash[key] = value
}

func main() {
	fmt.Println("Example 1")
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1)) // 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // -1
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // -1
	fmt.Println(lRUCache.Get(3)) // 3
	fmt.Println(lRUCache.Get(4)) // 4

	fmt.Println("Example 2")
	lRUCache = Constructor(2)
	lRUCache.Put(1, 0)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1)) // 0
	lRUCache.Put(3, 3)
	fmt.Println(lRUCache.Get(2)) // -1
	lRUCache.Put(4, 4)
	fmt.Println(lRUCache.Get(1))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(4))
}
