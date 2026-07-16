package _1331_1340

import "sort"

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 1331
func arrayRankTransform(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	sort.Ints(arrCopy)

	numToRank := make(map[int]int)
	rank := 1
	numToRank[arrCopy[0]] = rank

	for i := 1; i < len(arrCopy); i++ {
		if arrCopy[i] > arrCopy[i-1] {
			rank++
			numToRank[arrCopy[i]] = rank
		}
	}
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = numToRank[arr[i]]
	}

	return result
}

// leetcode problem No. 1339

func sumOfRootNode(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Val + sumOfRootNode(node.Left) + sumOfRootNode(node.Right)
}

func maxProduct(root *TreeNode) int {
	var recursive func(root *TreeNode) int
	maxVal := 0
	sumOfRoot := sumOfRootNode(root)
	recursive = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := recursive(node.Left)
		if product := leftSum * (sumOfRoot - leftSum); product > maxVal {
			maxVal = product
		}
		rightSum := recursive(node.Right)
		if product := rightSum * (sumOfRoot - rightSum); product > maxVal {
			maxVal = product
		}
		return node.Val + leftSum + rightSum
	}
	recursive(root)
	return maxVal%1_000_000_000 + 7
}
