package _1521_1530

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 1530

func countPairs(root *TreeNode, distance int) int {
	ans := 0
	var recursive func(*TreeNode) map[int]int /*return is the leaves count by depth*/
	recursive = func(node *TreeNode) map[int]int {
		if node.Left == nil && node.Right == nil {
			return map[int]int{
				1: 1,
			}
		}
		var leftLeavesCountByDepth map[int]int
		if node.Left != nil {
			leftLeavesCountByDepth = recursive(node.Left)
		}
		var rightLeavesCountByDepth map[int]int
		if node.Right != nil {
			rightLeavesCountByDepth = recursive(node.Right)
		}

		res := map[int]int{}
		if leftLeavesCountByDepth != nil && rightLeavesCountByDepth != nil {
			for depth, count := range leftLeavesCountByDepth {
				for i := 0; i <= distance-depth; i++ {
					ans += count * rightLeavesCountByDepth[i]
				}
			}

			// increase depth
			for depth, count := range leftLeavesCountByDepth {
				res[depth+1] += count
			}
			for depth, count := range rightLeavesCountByDepth {
				res[depth+1] += count
			}
			return res
		}

		if leftLeavesCountByDepth != nil {
			// increase depth
			for depth, count := range leftLeavesCountByDepth {
				res[depth+1] = count
			}
			return res
		}

		// if rightLeavesCountByDepth != nil
		// increase depth
		for depth, count := range rightLeavesCountByDepth {
			res[depth+1] = count
		}
		return res
	}

	recursive(root)

	return ans
}
