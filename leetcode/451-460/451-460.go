package _451_460

// leetcode problem No. 460

type LFUCacheNode struct {
	key   int
	value int
	prev  *LFUCacheNode
	next  *LFUCacheNode
	freq  int // frequency of current node
}

type LRUList struct { // basically represent an LRU
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
	nodeByKey    map[int]*LFUCacheNode
	lruByFreq    map[int] /*frequency*/ *LRUList
	capacity     int
	minFrequency int // this is used to fast local the lru list to delete node in put
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		nodeByKey:    make(map[int]*LFUCacheNode),
		lruByFreq:    make(map[int]*LRUList),
		capacity:     capacity,
		minFrequency: 0,
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.nodeByKey[key]; ok {
		lruCache := this.lruByFreq[node.freq]
		lruCache.remove(node)
		if lruCache.size == 0 && this.minFrequency == node.freq { // NOTE: this.minUseCounter == node.useCounter is necessary
			this.minFrequency++
		}

		node.freq++
		lruCache = this.lruByFreq[node.freq]
		if lruCache == nil {
			lruCache = makeLRUList()
		}
		lruCache.addToHead(node)
		this.lruByFreq[node.freq] = lruCache
		return node.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.nodeByKey[key]; ok {
		node.value = value
		this.Get(key)
		return
	}

	if len(this.nodeByKey) == this.capacity {
		lruCache := this.lruByFreq[this.minFrequency]
		nodeToRemove := lruCache.tail.prev
		lruCache.remove(nodeToRemove)
		delete(this.nodeByKey, nodeToRemove.key)
	}

	newNode := &LFUCacheNode{key: key, value: value, freq: 1}
	lruCache := this.lruByFreq[1]
	if lruCache == nil {
		lruCache = makeLRUList()
	}
	lruCache.addToHead(newNode)
	this.lruByFreq[1] = lruCache
	this.nodeByKey[key] = newNode
	this.minFrequency = 1
}
