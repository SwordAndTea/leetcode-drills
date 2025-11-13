package _51_160

import (
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	var solve func(index int) string
	solve = func(index int) string {
		if index == n {
			return ""
		}
		i := index
		for i < n && s[i] == ' ' {
			i++
		}
		j := i
		for j < n && s[j] != ' ' {
			j++
		}
		theWorld := s[i:j]
		otherWorld := solve(j)
		if otherWorld != "" {
			return strings.Join([]string{otherWorld, theWorld}, " ")
		}
		return theWorld
	}

	return solve(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	result := nums[0]
	// maxV/minV stores the max/min product of
	// subarray that ends with nums[i]
	maxV := nums[0]
	minV := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			// multiplied by a negative makes big number smaller, small number bigger
			tmp := maxV
			maxV = max(nums[i], minV*nums[i])
			minV = min(nums[i], tmp*nums[i])
		} else {
			maxV = max(nums[i], maxV*nums[i])
			minV = min(nums[i], minV*nums[i])
		}

		result = max(result, maxV)
	}
	return result
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if mid > 0 && nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if mid < n-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid] < nums[0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return nums[0]
}

func findMin2(nums []int) int {
	n := len(nums)
	var searchPivot func(left, right int) int
	searchPivot = func(left, right int) int {
		for left <= right {
			mid := (left + right) / 2
			if mid > 0 && nums[mid-1] > nums[mid] {
				return nums[mid]
			}
			if mid < n-1 && nums[mid] > nums[mid+1] {
				return nums[mid+1]
			}
			if nums[mid] > nums[n-1] || nums[mid] > nums[0] {
				left = mid + 1
			} else if nums[mid] < nums[n-1] || nums[mid] < nums[0] {
				right = mid - 1
			} else {
				// search left part
				p1 := searchPivot(left, mid-1)
				if p1 != -1 {
					return p1
				}
				// search right part
				p2 := searchPivot(mid+1, right)
				if p2 != -1 {
					return p2
				}
				return -1
			}
		}

		return -1
	}
	if pivot := searchPivot(0, n-1); pivot != -1 {
		return nums[pivot]
	}
	return nums[0]
}

type StackNode struct {
	Val  int
	Min  int
	Next *StackNode
}

type MinStack struct {
	Head *StackNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.Head == nil {
		this.Head = &StackNode{Val: val, Min: val}
	} else {
		this.Head = &StackNode{Val: val, Min: min(val, this.Head.Min), Next: this.Head}
	}
}

func (this *MinStack) Pop() {
	this.Head = this.Head.Next
}

func (this *MinStack) Top() int {
	return this.Head.Val
}

func (this *MinStack) GetMin() int {
	return this.Head.Min
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	count := 0
	pNeedForward := p2
	isP1Forward := false
	if p1 != nil {
		pNeedForward = p1
		isP1Forward = true
	}

	for pNeedForward != nil {
		pNeedForward = pNeedForward.Next
		count++
	}

	p1, p2 = headA, headB
	if isP1Forward {
		for count > 0 {
			p1 = p1.Next
			count--
		}
	} else {
		for count > 0 {
			p2 = p2.Next
			count--
		}
	}

	for p1 != nil && p2 != nil && p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	var searchPeakIndex func(left, right int) int
	searchPeakIndex = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + (right-left)/2
		if mid == 0 && nums[1] < nums[mid] {
			return mid
		}
		if mid == n-1 && nums[mid-1] < nums[mid] {
			return mid
		}
		if mid-1 >= 0 && nums[mid-1] < nums[mid] && mid+1 <= n-1 && nums[mid+1] < nums[mid] {
			return mid
		}
		if v := searchPeakIndex(left, mid-1); v != -1 {
			return v
		}
		if v := searchPeakIndex(mid+1, right); v != -1 {
			return v
		}
		return -1
	}

	return nums[searchPeakIndex(0, n-1)]
}
