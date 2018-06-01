/*
LRU Cache
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

Follow up:
Could you do both operations in O(1) time complexity?

Example:
LRUCache cache = new LRUCache( 2 );
cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.put(4, 4);    // evicts key 1
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4
*/

package main

// map + doubly list
type LRUCache struct {
	// memory
	memory 		map[int]int
	// pointer in pool
	pointer		map[int]*list.Element
	// select to evicts
	pool		*list.List
	// length
	len			int
	// capacity
	capacity 	int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		memory: 	make(map[int]int, capacity),
		pointer:	make(map[int]*list.Element, capacity),
		pool:		list.New(),
		len:		0,
		capacity: 	capacity,
	}
}

func (Cache *LRUCache) Get(key int) int {
	// key -> value
	if value,ok := Cache.memory[key];ok {
		Cache.pool.Remove(Cache.pointer[key])
		Cache.pointer[key] = Cache.pool.PushBack(key)
		return value
	}
	// invalid key
	return -1
}

func (Cache *LRUCache) Put(key int, value int) {
	if _,ok := Cache.memory[key];ok {
		Cache.pool.Remove(Cache.pointer[key])
	} else if Cache.len == Cache.capacity {
		invalid := Cache.pool.Remove(Cache.pool.Front()).(int)
		delete(Cache.memory, invalid)
		delete(Cache.pointer, invalid)
	} else {
		Cache.len++
	}
	// set key -> value
	Cache.memory[key] = value
	Cache.pointer[key] = Cache.pool.PushBack(key)
}