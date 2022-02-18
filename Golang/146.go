package main

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{key: key, val: value}
		this.cache[key] = node
		this.Add(node)
	}
	if len(this.cache) > this.capacity {
		delete(this.cache, key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
		node.next = nil
		return
	}
	if node == this.tail {
		this.tail = node.prev
	}
}

func (this *LRUCache) Add(node *Node) {
	node.prev = nil
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.next = nil
	}
}
