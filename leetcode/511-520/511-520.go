package _511_520

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	q := []*TreeNode{root}
	result := make([]int, 0)
	for len(q) > 0 {
		maxV := q[0].Val
		for i := 1; i < len(q); i++ {
			maxV = max(maxV, q[i].Val)
		}
		result = append(result, maxV)
		for i := len(q); i > 0; i-- {
			curNode := q[0]
			q = q[1:]
			if curNode.Left != nil {
				q = append(q, curNode.Left)
			}
			if curNode.Right != nil {
				q = append(q, curNode.Right)
			}
		}
	}
	return result
}
