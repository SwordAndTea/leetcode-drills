package _141_150

import (
	"strconv"
	"unsafe"
)

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodeInfo := map[int32]bool{}
	for head != nil {
		headAddr := *((*int32)(unsafe.Pointer(&head)))
		nodeInfo[headAddr] = true
		if head.Next != nil {
			nextAddr := *((*int32)(unsafe.Pointer(&head.Next)))
			if nodeInfo[nextAddr] {
				return true
			}
		}
		head = head.Next
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	nodeInfo := map[int32]*ListNode{}
	for head != nil {
		headAddr := *((*int32)(unsafe.Pointer(&head)))
		nodeInfo[headAddr] = head
		if head.Next != nil {
			nextAddr := *((*int32)(unsafe.Pointer(&head.Next)))
			if v, ok := nodeInfo[nextAddr]; ok {
				return v
			}
		}
		head = head.Next
	}
	return nil
}

func reorderList(head *ListNode) {
	nodeList := map[int]*ListNode{}
	i := 0
	for head != nil {
		nodeList[i] = head
		head = head.Next
		i++
	}
	i = 0
	j := len(nodeList) - 1
	p := &ListNode{0, nil}
	for i < j {
		p.Next = nodeList[i]
		p.Next.Next = nodeList[j]
		p = p.Next.Next
		i++
		j--
	}
	if i == j {
		p.Next = nodeList[i]
		p = p.Next
	}
	p.Next = nil
}

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node != nil {
			result = append(result, node.Val)
			preorder(node.Left)
			preorder(node.Right)
		}
	}
	preorder(root)
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node != nil {
			postorder(node.Left)
			postorder(node.Right)
			result = append(result, node.Val)
		}
	}
	postorder(root)
	return result
}

// leetcode problem No. 146

type LRUCache struct {
	Capacity int
	Cache    map[int]*LRUCacheNode
	Head     *LRUCacheNode
	Tail     *LRUCacheNode
}

// LRUCacheNode the bidirectional linked-list node of lru Cache
type LRUCacheNode struct {
	Key   int
	Value int
	Prev  *LRUCacheNode
	Next  *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Capacity: capacity,
		Cache:    make(map[int]*LRUCacheNode), // don't pass capacity to the make
		Head:     &LRUCacheNode{},
		Tail:     &LRUCacheNode{},
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
}

func (this *LRUCache) remove(node *LRUCacheNode) {
	nodePrev := node.Prev
	nodeNext := node.Next
	nodePrev.Next = nodeNext
	nodeNext.Prev = nodePrev
}

func (this *LRUCache) addToHead(node *LRUCacheNode) {
	headNext := this.Head.Next
	node.Next = headNext
	node.Prev = this.Head
	this.Head.Next = node
	headNext.Prev = node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.Cache[key]; ok {
		this.remove(v)
		this.addToHead(v)
		return v.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.Cache[key]; ok {
		v.Value = value
		this.remove(v)
		this.addToHead(v)
	} else {
		if len(this.Cache) == this.Capacity {
			nodeToRemove := this.Tail.Prev
			this.remove(nodeToRemove)
			delete(this.Cache, nodeToRemove.Key)
		}
		newNode := &LRUCacheNode{Key: key, Value: value}
		this.addToHead(newNode)
		this.Cache[key] = newNode
	}
}

func insertionSortList(head *ListNode) *ListNode {
	tmpHead := &ListNode{
		Val:  0,
		Next: head,
	}
	nodeToInsert := head.Next
	end := head
	for nodeToInsert != nil {
		p := tmpHead
		for p.Next != nodeToInsert && p.Next.Val <= nodeToInsert.Val {
			p = p.Next
		}
		if p == end {
			end = nodeToInsert
			nodeToInsert = nodeToInsert.Next
		} else {
			end.Next = nodeToInsert.Next
			nodeToInsert.Next = p.Next
			p.Next = nodeToInsert
			nodeToInsert = end.Next
		}
	}
	return tmpHead.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmpHead := &ListNode{Next: head}
	p1, p2 := tmpHead, tmpHead
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	l2 := sortList(p1.Next)
	p1.Next = nil
	l1 := sortList(head)

	// merge
	p := tmpHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	for l1 != nil {
		p.Next = l1
		l1 = l1.Next
		p = p.Next
	}

	for l2 != nil {
		p.Next = l2
		l2 = l2.Next
		p = p.Next
	}
	return tmpHead.Next
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}

	result := 0
	for i := 0; i < n; i++ {
		line := make(map[int] /*delta x*/ map[int] /*delta y*/ int) // the max points in the line with a slope of delta x/ delta y
		overlap := 0
		max := 0
		for j := i + 1; j < n; j++ {
			// calculate the slope of points[j] and points[i]
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			if x == 0 && y == 0 {
				overlap++
				continue
			}
			// simplify the slope
			theGCD := gcd(x, y)
			if theGCD != 0 {
				x /= theGCD
				y /= theGCD
			}

			// get max
			if v1, ok1 := line[x]; ok1 {
				if v2, ok2 := v1[y]; ok2 {
					v1[y] = v2 + 1
				} else {
					v1[y] = 1
				}
			} else {
				v1 = make(map[int]int)
				v1[y] = 1
				line[x] = v1
			}

			if line[x][y] > max {
				max = line[x][y]
			}
		}
		if v := max + overlap + 1; v > result {
			result = v
		}
	}
	return result
}

func evalRPN(tokens []string) int {
	numStack := make([]int, len(tokens))
	top := 0
	for _, token := range tokens {
		if token == "+" {
			numStack[top-2] = numStack[top-2] + numStack[top-1]
			top--
		} else if token == "-" {
			numStack[top-2] = numStack[top-2] - numStack[top-1]
			top--
		} else if token == "*" {
			numStack[top-2] = numStack[top-2] * numStack[top-1]
			top--
		} else if token == "/" {
			numStack[top-2] = numStack[top-2] / numStack[top-1]
			top--
		} else {
			num, _ := strconv.Atoi(token)
			numStack[top] = num
			top++
		}
	}
	return numStack[0]
}
