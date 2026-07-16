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
			return map[int]int{1: 1}
		}

		var leftLeafNodeByDistance map[int]int
		if node.Left != nil {
			leftLeafNodeByDistance = recursive(node.Left)
		}

		var rightLeafNodeByDistance map[int]int
		if node.Right != nil {
			rightLeafNodeByDistance = recursive(node.Right)
		}

		if leftLeafNodeByDistance != nil && rightLeafNodeByDistance != nil {
			for depth, numLeafNodes := range leftLeafNodeByDistance {
				for i := 1; i <= distance-depth; i++ {
					ans += numLeafNodes * rightLeafNodeByDistance[i]
				}
			}
		}

		// increase depth and merge result
		newInfo := make(map[int]int)
		for depth, numLeafNodes := range leftLeafNodeByDistance {
			newInfo[depth+1] += numLeafNodes
		}

		for depth, numLeafNodes := range rightLeafNodeByDistance {
			newInfo[depth+1] += numLeafNodes
		}

		return newInfo
	}

	recursive(root)
	return ans
}
