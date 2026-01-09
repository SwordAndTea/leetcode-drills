package _861_870

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTheSubtreeByCalDepth(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	leftDepth, leftDeepestSubtree := getTheSubtreeByCalDepth(root.Left)
	rightDepth, rightDeepestSubtree := getTheSubtreeByCalDepth(root.Right)
	if leftDepth == rightDepth {
		return 1 + leftDepth, root
	}

	if leftDepth > rightDepth {
		return 1 + leftDepth, leftDeepestSubtree
	}

	return 1 + rightDepth, rightDeepestSubtree
}

// leetcode problem No. 865
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, res := getTheSubtreeByCalDepth(root)
	return res
}
