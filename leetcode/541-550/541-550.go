package _541_550

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leftBoundary(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	if node.Left != nil {
		nextBoundary := leftBoundary(node.Left)
		return append([]int{node.Val}, nextBoundary...)
	}
	nextBoundary := leftBoundary(node.Right)
	return append([]int{node.Val}, nextBoundary...)
}

func rightBoundary(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	if node.Right != nil {
		nextBoundary := rightBoundary(node.Right)
		return append([]int{node.Val}, nextBoundary...)
	}
	nextBoundary := rightBoundary(node.Left)
	return append([]int{node.Val}, nextBoundary...)
}

func leaves(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	var leftLeaves []int
	var rightLeaves []int
	if node.Left != nil {
		leftLeaves = leaves(node.Left)
	}
	if node.Right != nil {
		rightLeaves = leaves(node.Right)
	}
	return slices.Concat(leftLeaves, rightLeaves)
}

// leetcode problem No. 545
func boundaryOfBinaryTree(root *TreeNode) []int {
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	rightB := rightBoundary(root.Right)
	slices.Reverse(rightB)
	return slices.Concat([]int{root.Val}, leftBoundary(root), leaves(root), rightB)
}
