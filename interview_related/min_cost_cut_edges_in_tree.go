package interview_related

// question description: https://leetcode.com/discuss/post/6490748/google-interview-question-help-onsite-1-drwvu/

type TreeNode struct {
	Val   int // the edge cost from cur node to its parent
	Left  *TreeNode
	Right *TreeNode
}

func minCostCut(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
		leftCost := 0
		if node.Left != nil {
			leftCost = recursive(node.Left)
		}
		rightCost := 0
		if node.Right != nil {
			rightCost = recursive(node.Right)
		}
		return min(node.Val, leftCost+rightCost)
	}

	return recursive(root.Left) + recursive(root.Right)
}

func minCostCutWithNegativeCost(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
		leftCost := 0
		if node.Left != nil {
			leftCost = recursive(node.Left)
		}
		rightCost := 0
		if node.Right != nil {
			rightCost = recursive(node.Right)
		}
		return min(node.Val, leftCost+rightCost, node.Val+leftCost+rightCost, node.Val+leftCost, node.Val+rightCost)
	}

	return recursive(root.Left) + recursive(root.Right)
}
