package _2501_2510

// leetcode problem No. 2502

type MemoryBlock struct {
	index int
	size  int
	prev  *MemoryBlock
	next  *MemoryBlock
}

type Allocator struct {
	head     *MemoryBlock
	tail     *MemoryBlock
	blockMap map[int][]*MemoryBlock
	cap      int
	curSize  int
}

func Constructor(n int) Allocator {
	head := &MemoryBlock{}
	tail := &MemoryBlock{
		index: n,
		size:  0,
	}
	head.next = tail
	tail.prev = head
	return Allocator{
		head:     head,
		tail:     tail,
		blockMap: make(map[int][]*MemoryBlock),
		cap:      n,
	}
}

func (this *Allocator) insertAfter(node *MemoryBlock, after *MemoryBlock) {
	afterNext := after.next
	afterNext.prev = node
	node.next = afterNext
	node.prev = after
	after.next = node
}

func (this *Allocator) delete(node *MemoryBlock) {
	preNode := node.prev
	afterNode := node.next
	node.prev = nil
	node.next = nil
	preNode.next = afterNode
	afterNode.prev = preNode
}

func (this *Allocator) Allocate(size int, mID int) int {
	if this.curSize+size > this.cap {
		return -1
	}
	cur := this.head.next
	startIndex := 0
	for cur != nil {
		if cur.index-startIndex >= size { // do insert
			block := &MemoryBlock{
				index: startIndex,
				size:  size,
			}
			this.insertAfter(block, cur.prev)
			this.blockMap[mID] = append(this.blockMap[mID], block)
			this.curSize += size
			return startIndex
		}
		startIndex = cur.index + cur.size
		cur = cur.next
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	ans := 0

	for _, block := range this.blockMap[mID] {
		ans += block.size
		this.delete(block)
	}

	this.blockMap[mID] = make([]*MemoryBlock, 0)
	this.curSize -= ans

	return ans
}
