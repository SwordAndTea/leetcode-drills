package _451_460

// leetcode problem No. 460

type LFUCacheNode struct {
	key        int
	value      int
	prev       *LFUCacheNode
	next       *LFUCacheNode
	useCounter int
}

type LRUList struct { // basically represent a LRU
	head *LFUCacheNode
	tail *LFUCacheNode
	size int
}

func makeLRUList() *LRUList {
	head := &LFUCacheNode{}
	tail := &LFUCacheNode{}
	head.next = tail
	tail.prev = head
	return &LRUList{head: head, tail: tail, size: 0}
}

func (lru *LRUList) remove(node *LFUCacheNode) {
	nodePrev := node.prev
	nodeNext := node.next
	nodePrev.next = nodeNext
	nodeNext.prev = nodePrev
	lru.size--
}

func (lru *LRUList) addToHead(node *LFUCacheNode) {
	headNext := lru.head.next
	lru.head.next = node
	headNext.prev = node
	node.prev = lru.head
	node.next = headNext
	lru.size++
}

type LFUCache struct {
	nodeMap       map[int]*LFUCacheNode
	lruMap        map[int] /*use counter*/ *LRUList
	capacity      int
	minUseCounter int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodeMap:       make(map[int]*LFUCacheNode),
		lruMap:        make(map[int]*LRUList),
		capacity:      capacity,
		minUseCounter: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.nodeMap[key]; ok {
		lruCache := this.lruMap[node.useCounter]
		lruCache.remove(node)
		if lruCache.size == 0 && this.minUseCounter == node.useCounter { // NOTE: this.minUseCounter == node.useCounter is necessary
			this.minUseCounter++
		}

		node.useCounter++
		lruCache = this.lruMap[node.useCounter]
		if lruCache == nil {
			lruCache = makeLRUList()
		}
		lruCache.addToHead(node)
		this.lruMap[node.useCounter] = lruCache
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.nodeMap[key]; ok {
		node.value = value
		this.Get(key)
		return
	}

	if len(this.nodeMap) == this.capacity {
		lruCache := this.lruMap[this.minUseCounter]
		nodeToRemove := lruCache.tail.prev
		lruCache.remove(nodeToRemove)
		delete(this.nodeMap, nodeToRemove.key)
	}

	newNode := &LFUCacheNode{key: key, value: value, useCounter: 1}
	lruCache := this.lruMap[1]
	if lruCache == nil {
		lruCache = makeLRUList()
	}
	lruCache.addToHead(newNode)
	this.lruMap[1] = lruCache
	this.nodeMap[key] = newNode
	this.minUseCounter = 1
}
