package _1161_1170

import "math"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 1161
func maxLevelSum(root *TreeNode) int {
	maxVal := math.MinInt
	maxValLevel := 1
	if root == nil {
		return maxValLevel
	}

	queue := []*TreeNode{root}
	curIndex := 0
	curLevel := 1
	curLevelEnd := 0
	nextLevelEnd := 0
	curLevelSum := 0
	for curIndex < len(queue) {
		curNode := queue[curIndex]
		curLevelSum += curNode.Val
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			nextLevelEnd++
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			nextLevelEnd++
		}
		if curIndex == curLevelEnd {
			if curLevelSum > maxVal {
				maxVal = curLevelSum
				maxValLevel = curLevel
			}
			curLevel++
			curLevelEnd = nextLevelEnd
			curLevelSum = 0
		}
		curIndex++
	}

	return maxValLevel
}
