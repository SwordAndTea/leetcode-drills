package _231_240

func isPowerOfTwo(n int) bool {
	result := 0
	for n != 0 {
		n &= n - 1
		result++
	}
	return result == 1
}

type MyQueue struct {
	data  []int
	front int
	rear  int
}

func Constructor() MyQueue {
	return MyQueue{data: make([]int, 0, 128), front: 0, rear: 0}
}

func (this *MyQueue) Push(x int) {
	if this.rear == len(this.data) {
		this.data = append(this.data, x)
	} else {
		this.data[this.rear] = x
	}
	this.rear++
}

func (this *MyQueue) Pop() int {
	v := this.data[this.front]
	this.front++
	return v
}

func (this *MyQueue) Peek() int {
	return this.data[this.front]
}

func (this *MyQueue) Empty() bool {
	return this.front == this.rear
}

func countDigitOne(n int) int {
	ans := 0
	for i := 1; i <= n; i *= 10 {
		prefix := n / (i * 10)
		digit := n / i % 10
		suffix := n % i

		if digit == 0 {
			ans += prefix * i
		} else if digit == 1 {
			ans += prefix*i + suffix + 1
		} else { // digit is in [2, 9]
			ans += (prefix + 1) * i
		}
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode problem No. 234
func isPalindrome(head *ListNode) bool {
	midP := head
	p := head
	for p.Next != nil && p.Next.Next != nil {
		p = p.Next.Next
		midP = midP.Next
	}

	// do head insert at mid
	p = midP.Next
	midP.Next = nil // necessary step to make sure after head insert, the next pointer of the original node that after mid is pointed to nil
	for p != nil {
		nextP := p.Next
		p.Next = midP.Next
		midP.Next = p
		p = nextP
	}

	p1 := head
	p2 := midP.Next
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 235, Lowest Common Ancestor of a Binary Search Tree
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val == root.Val {
		return p
	}
	if q.Val == root.Val {
		return q
	}

	if (p.Val < root.Val && q.Val > root.Val) || (p.Val > root.Val && q.Val < root.Val) {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return lowestCommonAncestor(root.Right, p, q)
}

// leetcode problem No. 236
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { // if curNode is p or q
		return root
	}
	l := lowestCommonAncestor2(root.Left, p, q)  // try to find p or q in left child
	r := lowestCommonAncestor2(root.Right, p, q) // try to find p or q in right child
	// * if we find p in left child and also find q in left child, the r will be nil
	// and the result will be the left child or the children of left child

	// * if we find p in right child and also find q in right child, the l will be nil
	// and the result will be the right child or the children of right child

	// * if we find p or q in left child and also find q or q in right child, the p and q
	// must be located in two separate children, then the root node is the actual common ancestor
	if l != nil && r != nil {
		return root
	}
	if l != nil {
		return l
	}
	if r != nil {
		return r
	}
	return nil
}

func deleteNode(node *ListNode) {
	p1 := node
	p2 := node.Next
	for p2 != nil {
		p1.Val = p2.Val
		if p2.Next == nil {
			p1.Next = nil
		} else {
			p1 = p1.Next
		}
		p2 = p2.Next
	}
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefixProduct := make([]int, n)
	prefixProduct[0] = nums[0]
	suffixProduct := make([]int, n)
	suffixProduct[n-1] = nums[n-1]
	for i := 1; i < n-1; i++ {
		prefixProduct[i] = prefixProduct[i-1] * nums[i]
	}
	for i := n - 2; i >= 1; i-- {
		suffixProduct[i] = suffixProduct[i+1] * nums[i]
	}
	result := make([]int, n)
	result[0] = suffixProduct[1]
	result[n-1] = prefixProduct[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = prefixProduct[i-1] * suffixProduct[i+1]
	}
	return result
}
