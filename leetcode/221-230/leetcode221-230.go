package _221_230

import (
	"fmt"
	"math"
)

// leetcode problem No. 221
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	heights := make([]int, n)
	left := make([]int, n)
	right := make([]int, n)
	for j := 0; j < n; j++ {
		right[j] = n
	}

	theMax := 0
	for i := 0; i < m; i++ {
		leftest := 0
		rightest := n - 1
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++

				if leftest > left[j] {
					left[j] = leftest
				}
			} else {
				heights[j] = 0

				left[j] = 0
				leftest = j + 1
			}

			if matrix[i][n-1-j] == '1' {
				if rightest < right[n-1-j] {
					right[n-1-j] = rightest
				}
			} else {
				right[n-1-j] = n
				rightest = n - 1 - j - 1
			}

		}

		area := 0
		for j := 0; j < n; j++ {
			if v := right[j] - left[j] + 1; v < heights[j] {
				area = v * v
			} else {
				area = heights[j] * heights[j]
			}
			if area > theMax {
				theMax = area
			}
		}
	}

	return theMax
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + treeDepth(root.Left)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := treeDepth(root.Left)
	rightDepth := treeDepth(root.Right)
	if leftDepth == rightDepth { // this means left subtree is a full binary tree
		return int(math.Pow(2, float64(leftDepth))) + countNodes(root.Right)
	} else { // right subtree is a full binary tree
		return int(math.Pow(2, float64(rightDepth))) + countNodes(root.Left)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int,
	bx1 int, by1 int, bx2 int, by2 int) int {
	innerAreaTopLeftX := max(ax1, bx1)
	innerAreaTopLeftY := min(ay2, by2)

	innerAreaBottomRightX := min(ax2, bx2)
	innerAreaBottomRightY := max(ay1, by1)

	innerWidth := max(0, innerAreaBottomRightX-innerAreaTopLeftX)
	innerHeight := max(0, innerAreaTopLeftY-innerAreaBottomRightY)

	innerArea := innerWidth * innerHeight

	unionArea := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	return unionArea - innerArea
}

type MyStack struct {
	values []int
	top    int
}

func Constructor() MyStack {
	return MyStack{
		values: []int{},
		top:    0,
	}
}

func (this *MyStack) Push(x int) {
	if len(this.values) == this.top {
		this.values = append(this.values, x)
	} else {
		this.values[this.top] = x
	}
	this.top++
}

func (this *MyStack) Pop() int {
	v := this.values[this.top-1]
	this.top--
	return v
}

func (this *MyStack) Top() int {
	return this.values[this.top-1]
}

func (this *MyStack) Empty() bool {
	return this.top == 0
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func summaryRanges(nums []int) []string {
	n := len(nums)
	result := make([]string, 0)
	startIndex := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		if i-1 == startIndex {
			result = append(result, fmt.Sprint(nums[startIndex]))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", nums[startIndex], nums[i-1]))
		}
		startIndex = i
	}
	if startIndex == n-1 {
		result = append(result, fmt.Sprint(nums[startIndex]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[startIndex], nums[n-1]))
	}
	return result
}

func majorityElement(nums []int) []int {
	n := len(nums)
	if n < 1 {
		return nums
	}
	m1, c1 := nums[0], 1
	m2, c2 := math.MaxInt, 0
	for i := 1; i < n; i++ {
		if nums[i] == m1 {
			c1++
			if c2 > 0 {
				c2--
			}
		} else if nums[i] == m2 {
			if c1 > 0 {
				c1--
			}
			c2++
		} else {
			if c1 == 0 {
				m1 = nums[i]
				c1 = 1
				continue
			}
			if c2 == 0 {
				m2 = nums[i]
				c2 = 1
				continue
			}
			c1--
			c2--
		}
	}

	result := make([]int, 0, 2)
	realCount1, realCount2 := 0, 0
	for _, num := range nums {
		if c1 != 0 && num == m1 {
			realCount1++
		}
		if c2 != 0 && num == m2 {
			realCount2++
		}
	}
	if realCount1 > n/3 {
		result = append(result, m1)
	}
	if realCount2 > n/3 {
		result = append(result, m2)
	}
	return result
}

func kthSmallest(root *TreeNode, k int) int {
	var solve func(node *TreeNode, startIndex int) (int, int)

	solve = func(node *TreeNode, startIndex int) (int, int) {
		if node == nil {
			return 0, -1
		}

		leftChildrenCount, v1 := solve(node.Left, startIndex)
		if v1 != -1 {
			return 0, v1
		}
		curIndex := startIndex + leftChildrenCount + 1
		if curIndex == k {
			return curIndex, node.Val
		}
		rightChildrenCount, v2 := solve(node.Right, curIndex)
		if v2 != -1 {
			return 0, v2
		}
		return leftChildrenCount + rightChildrenCount + 1, -1
	}

	_, v := solve(root, 0)
	return v
}
