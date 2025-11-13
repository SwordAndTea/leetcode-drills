package _01_110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	//nodes := make([]*TreeNode, 0, 4)
	//nodes = append(nodes, root)
	//curP, topP := 0, 0
	//for curP <= topP {
	//	p1, p2 := curP, topP
	//	for p1 < p2 {
	//		if nodes[p1].Left == nil && nodes[p2].Right == nil {
	//			// do nothing
	//		} else if nodes[p1].Left == nil || nodes[p2].Right == nil {
	//			return false
	//		} else {
	//			if nodes[p1].Left.Val != nodes[p2].Right.Val {
	//				return false
	//			}
	//		}
	//
	//		if nodes[p1].Right == nil && nodes[p2].Left == nil {
	//			// do nothing
	//		} else if nodes[p1].Right == nil || nodes[p2].Left == nil {
	//			return false
	//		} else {
	//			if nodes[p1].Right.Val != nodes[p2].Left.Val {
	//				return false
	//			}
	//		}
	//
	//		p1++
	//		p2--
	//	}
	//	for curP <= topP {
	//		if nodes[curP].Left != nil {
	//			nodes = append(nodes, nodes[curP].Left)
	//		}
	//		if nodes[curP].Right != nil {
	//			nodes = append(nodes, nodes[curP].Right)
	//		}
	//		curP++
	//	}
	//	topP = len(nodes) - 1
	//}
	//return true

	var judge func(p1 *TreeNode, p2 *TreeNode) bool

	judge = func(p1 *TreeNode, p2 *TreeNode) bool {
		if p1 == nil && p2 == nil {
			return true
		}

		if p1 == nil || p2 == nil {
			return false
		}

		if p1.Val != p2.Val {
			return false
		}

		return judge(p1.Left, p2.Right) && judge(p1.Right, p2.Left)
	}

	return judge(root.Left, root.Right)
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	nodes := make([]*TreeNode, 0, 4)
	nodes = append(nodes, root)
	curP, topP := 0, 0
	for curP <= topP {
		levelNodes := make([]int, 0, topP-curP+1)
		for curP <= topP {
			levelNodes = append(levelNodes, nodes[curP].Val)
			if nodes[curP].Left != nil {
				nodes = append(nodes, nodes[curP].Left)
			}
			if nodes[curP].Right != nil {
				nodes = append(nodes, nodes[curP].Right)
			}
			curP++
		}
		result = append(result, levelNodes)
		topP = len(nodes) - 1
	}
	return result
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	nodes := make([]*TreeNode, 0, 4)
	nodes = append(nodes, root)
	curP, topP := 0, 0
	leftToRight := true
	for curP <= topP {
		levelNodes := make([]int, 0, topP-curP+1)
		p1, p2 := curP, topP
		for p2 >= p1 {
			levelNodes = append(levelNodes, nodes[p2].Val)
			if leftToRight {
				if nodes[p2].Left != nil {
					nodes = append(nodes, nodes[p2].Left)
				}
				if nodes[p2].Right != nil {
					nodes = append(nodes, nodes[p2].Right)
				}
			} else {
				if nodes[p2].Right != nil {
					nodes = append(nodes, nodes[p2].Right)
				}
				if nodes[p2].Left != nil {
					nodes = append(nodes, nodes[p2].Left)
				}
			}

			p2--
		}
		result = append(result, levelNodes)
		curP = topP + 1
		topP = len(nodes) - 1
		leftToRight = !leftToRight
	}
	return result
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	rootIndex := -1
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])
	return root
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	rootIndex := -1
	for i, v := range inorder {
		if v == root.Val {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree2(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree2(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0, 4)
	var traversal func(levelNodes []*TreeNode)
	traversal = func(levelNodes []*TreeNode) {
		if len(levelNodes) == 0 {
			return
		}
		nextLevelNodes := make([]*TreeNode, 0, len(levelNodes))
		oneResult := make([]int, 0, len(levelNodes))
		for _, n := range levelNodes {
			oneResult = append(oneResult, n.Val)
			if n.Left != nil {
				nextLevelNodes = append(nextLevelNodes, n.Left)
			}
			if n.Right != nil {
				nextLevelNodes = append(nextLevelNodes, n.Right)
			}
		}
		traversal(nextLevelNodes)
		result = append(result, oneResult)
	}
	traversal([]*TreeNode{root})
	return result
}

func sortedArrayToBST(nums []int) *TreeNode {
	var solve func(remain []int) *TreeNode

	solve = func(remain []int) *TreeNode {
		if len(remain) == 0 {
			return nil
		}
		mid := len(remain) / 2
		node := &TreeNode{
			Val:   remain[mid],
			Left:  solve(remain[0:mid]),
			Right: solve(remain[mid+1:]),
		}
		return node
	}

	return solve(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	values := make([]int, 0, 16)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(values)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var judge func(node *TreeNode) (bool, int)
	judge = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		leftBalanced, leftHeight := judge(node.Left)
		rightBalanced, rightHeight := judge(node.Right)
		return leftBalanced && rightBalanced && abs(leftHeight-rightHeight) <= 1, max(leftHeight, rightHeight) + 1
	}

	res, _ := judge(root)
	return res
}
