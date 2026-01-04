package _91_100

import "strings"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1) // num of decodings from (s[0], s[i])
	dp[0] = 1
	dp[1] = 1
	if s[0] == '0' {
		dp[1] = 0
	}

	for i := 2; i <= len(s); i++ {
		if s[i-1] == '0' {
			if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 0
			}
		} else if s[i-1] <= '6' {
			if s[i-2] == '1' || s[i-2] == '2' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		} else {
			if s[i-2] == '1' {
				dp[i] = dp[i-1] + dp[i-2]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}

	return dp[len(s)]
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

func isInterleave(s1 string, s2 string, s3 string) bool {
	//if len(s3) != len(s1)+len(s2) {
	//	return false
	//}
	//
	//if s3 == s1+s2 {
	//	return true
	//}
	//
	//memo := make(map[string]map[string]map[string]map[bool]bool)
	//
	//setMemo := func(s1Remain, s2Remain, s3Remain string, matchS1 bool, val bool) {
	//	m1 := memo[s1Remain]
	//	if m1 == nil {
	//		m1 = make(map[string]map[string]map[bool]bool)
	//	}
	//	memo[s1Remain] = m1
	//
	//	m2 := m1[s2Remain]
	//	if m2 == nil {
	//		m2 = make(map[string]map[bool]bool)
	//	}
	//	m1[s2Remain] = m2
	//
	//	m3 := m2[s3Remain]
	//	if m3 == nil {
	//		m3 = make(map[bool]bool)
	//	}
	//	m2[s3Remain] = m3
	//
	//	m3[matchS1] = val
	//}
	//
	//getMeme := func(s1Remain, s2Remain, s3Remain string, matchS1 bool) (bool, bool) {
	//	m1 := memo[s1Remain]
	//	if m1 == nil {
	//		m1 = make(map[string]map[string]map[bool]bool)
	//	}
	//	memo[s1Remain] = m1
	//
	//	m2 := m1[s2Remain]
	//	if m2 == nil {
	//		m2 = make(map[string]map[bool]bool)
	//	}
	//	m1[s2Remain] = m2
	//
	//	m3 := m2[s3Remain]
	//	if m3 == nil {
	//		m3 = make(map[bool]bool)
	//	}
	//	m2[s3Remain] = m3
	//	if val, ok := m3[matchS1]; ok {
	//		return val, ok
	//	} else {
	//		return val, ok
	//	}
	//}
	//
	//var solve func(s1Remain, s2Remain, s3Remain string, matchS1 bool) bool
	//solve = func(s1Remain, s2Remain, s3Remain string, matchS1 bool) bool {
	//	if len(s1Remain) == 0 && len(s2Remain) == 0 && len(s3Remain) == 0 {
	//		return true
	//	}
	//
	//	if val, ok := getMeme(s1Remain, s2Remain, s3Remain, matchS1); ok {
	//		return val
	//	}
	//	if matchS1 {
	//		for i := 0; i < len(s1Remain); i++ {
	//			if s1Remain[0:i+1] == s3Remain[0:i+1] {
	//				isMatch := solve(s1Remain[i+1:], s2Remain, s3Remain[i+1:], false)
	//				setMemo(s1Remain[i+1:], s2Remain, s3Remain[i+1:], false, isMatch)
	//				if isMatch {
	//					return true
	//				}
	//			} else {
	//				break
	//			}
	//		}
	//	} else {
	//		for i := 0; i < len(s2Remain); i++ {
	//			if s2Remain[0:i+1] == s3Remain[0:i+1] {
	//				isMatch := solve(s1Remain, s2Remain[i+1:], s3Remain[i+1:], true)
	//				setMemo(s1Remain, s2Remain[i+1:], s3Remain[i+1:], true, isMatch)
	//				if isMatch {
	//					return true
	//				}
	//			} else {
	//				break
	//			}
	//		}
	//	}
	//	return false
	//}
	//
	//if solve(s1, s2, s3, true) {
	//	return true
	//}
	//
	//return solve(s1, s2, s3, false)

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
