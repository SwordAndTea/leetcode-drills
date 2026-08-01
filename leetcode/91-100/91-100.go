package _91_100

import "strings"

// leetcode problem No. 91

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	if s[n-1] != '0' {
		dp[n-1] = 1
	}
	num := s[n-1] - '1' + 1
	for i := n - 2; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			num = 0
			continue
		}
		dp[i] = dp[i+1]
		num = (s[i]-'1'+1)*10 + num
		if num >= 1 && num <= 26 {
			dp[i] += dp[i+2]
		}
		num = num / 10
	}

	return dp[0]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head.Next == nil {
		return head
	}
	if left == right {
		return head
	}
	i := 1
	leftP := head
	for i < left {
		leftP = leftP.Next
		i++
	}

	stack := make([]int, right-left+1)
	top := 0
	rightP := leftP
	for i < right {
		stack[top] = rightP.Val
		top++
		rightP = rightP.Next
		i++
	}
	stack[top] = rightP.Val
	top++

	for i = left; i < right; i++ {
		top--
		leftP.Val = stack[top]
		leftP = leftP.Next

	}
	top--
	leftP.Val = stack[top]

	return head
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)

	var solve func(curI int, currentList []string)

	solve = func(curI int, currentList []string) {
		if len(currentList) == 4 && curI == len(s) {
			result = append(result, strings.Join(currentList, "."))
			return
		}

		if curI == len(s) || len(currentList) == 4 {
			return
		}

		if s[curI] == '0' {
			solve(curI+1, append(currentList, "0"))
		} else {
			value := 0
			for i := curI; i < len(s); i++ {
				value = value*10 + int(s[i]-'0')
				if value <= 255 {
					solve(i+1, append(currentList, s[curI:i+1]))
				} else {
					break
				}
			}
		}
	}

	solve(0, []string{})

	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var solve func(node *TreeNode)

	solve = func(node *TreeNode) {
		if node != nil {
			solve(node.Left)
			result = append(result, node.Val)
			solve(node.Right)
		}
	}

	solve(root)
	return result
}

func generateTrees(n int) []*TreeNode {
	var solve func(start, end int) []*TreeNode
	solve = func(start, end int) []*TreeNode {
		result := make([]*TreeNode, 0)
		for i := start; i <= end; i++ {
			leftTree := solve(start, i-1)
			rightTree := solve(i+1, end)

			if len(leftTree) == 0 && len(rightTree) == 0 {
				result = append(result, &TreeNode{
					Val:   i,
					Left:  nil,
					Right: nil,
				})
			} else if len(leftTree) != 0 && len(rightTree) != 0 {
				for _, v1 := range leftTree {
					for _, v2 := range rightTree {
						result = append(result, &TreeNode{
							Val:   i,
							Left:  v1,
							Right: v2,
						})
					}
				}
			} else if len(leftTree) != 0 {
				for _, v1 := range leftTree {
					result = append(result, &TreeNode{
						Val:   i,
						Left:  v1,
						Right: nil,
					})
				}
			} else {
				for _, v2 := range rightTree {
					result = append(result, &TreeNode{
						Val:   i,
						Left:  nil,
						Right: v2,
					})
				}
			}
		}

		return result
	}

	return solve(1, n)
}

func numTrees(n int) int {
	var solve func(start, end int) int
	memo := make(map[int]int)
	solve = func(start, end int) int {
		if v, ok := memo[end-start+1]; ok {
			return v
		}
		result := 0
		for i := start; i <= end; i++ {
			leftTree := solve(start, i-1)
			rightTree := solve(i+1, end)

			if leftTree == 0 && rightTree == 0 {
				result += 1
			} else if leftTree != 0 && rightTree != 0 {
				result += leftTree * rightTree
			} else if leftTree != 0 {
				result += leftTree
			} else {
				result += rightTree
			}
		}
		memo[end-start+1] = result
		return result
	}

	return solve(1, n)
}

// leetcode problem No. 97
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}

	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	m := len(s1)
	n := len(s2)
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = false

	for i := 1; i < m+1; i++ {
		dp[i][0] = s1[0:i] == s3[0:i]
	}

	for j := 1; j < n+1; j++ {
		dp[0][j] = s2[0:j] == s3[0:j]
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			// as dp[i][j] is only relevant to dp[i-1][j] and dp[i][j-1]
			// the space complexity can be reduced to o(len(s2))
			dp[i][j] = false
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	isLeftValid := true
	isRightValid := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		// find pre node
		p := root.Left
		for p.Right != nil {
			p = p.Left
		}
		if p.Val >= root.Val {
			return false
		}
		isLeftValid = isValidBST(root.Left)
	}
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		p := root.Right
		for p.Left != nil {
			p = p.Left
		}
		if p.Val <= root.Val {
			return false
		}
		isRightValid = isValidBST(root.Right)
	}
	return isLeftValid && isRightValid
}

func recoverTree(root *TreeNode) {
	nodeList := make([]*TreeNode, 0, 2)
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			nodeList = append(nodeList, node)
			inorder(node.Right)
		}
	}

	inorder(root)

	i := 0
	for i+1 < len(nodeList) && nodeList[i+1].Val > nodeList[i].Val {
		i++
	}
	j := len(nodeList) - 1
	for j-1 >= 0 && nodeList[j-1].Val < nodeList[j].Val {
		j--
	}
	nodeList[i].Val, nodeList[j].Val = nodeList[j].Val, nodeList[i].Val
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
