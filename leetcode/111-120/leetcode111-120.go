package _11_120

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil && root.Right != nil {
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
	if root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + minDepth(root.Right)
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var judge func(node *TreeNode, current int) bool
	judge = func(node *TreeNode, current int) bool {
		current += node.Val
		if node.Left == nil && node.Right == nil {
			return current == targetSum
		}

		if node.Left != nil && node.Right != nil {
			return judge(node.Left, current) || judge(node.Right, current)
		}

		if node.Left != nil {
			return judge(node.Left, current)
		}

		return judge(node.Right, current)
	}

	return judge(root, 0)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	var solve func(node *TreeNode, current []int, currentSum int)
	solve = func(node *TreeNode, current []int, currentSum int) {
		currentSum += node.Val
		current = append(current, node.Val)
		if node.Left == nil && node.Right == nil {
			if currentSum == targetSum {
				newResult := make([]int, len(current))
				copy(newResult, current)
				result = append(result, newResult)
			}
		}

		if node.Left != nil {
			solve(node.Left, current, currentSum)
		}

		if node.Right != nil {
			solve(node.Right, current, currentSum)
		}
	}

	solve(root, []int{}, 0)
	return result
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Right)

	if root.Left != nil {
		flatten(root.Left)
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	//for j := 1; j < n; j++ {
	//	dp[0][j] = 0
	//}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
		if s[i] == t[0] {
			dp[i][0] += 1
		}
		for j := 1; j < n && j <= i; j++ {
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	nodeList := make([]*Node, 1, 2)
	nodeList[0] = root
	top := 1
	cur := 0
	var pre *Node
	for cur < top {
		curNode := nodeList[cur]
		if pre != nil {
			pre.Next = curNode
		}
		pre = curNode
		if curNode.Left != nil {
			nodeList = append(nodeList, curNode.Left)
		}
		if curNode.Right != nil {
			nodeList = append(nodeList, curNode.Right)
		}
		cur += 1
		if cur == top {
			top = len(nodeList)
			pre = nil
		}
	}
	return root
}

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = make([]int, 1)
	result[0][0] = 1
	for i := 1; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = result[i-1][0]
		result[i][i] = result[i-1][i-1]
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
	}
	return result
}

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		result[i] = result[i-1]
		for j := i - 1; j > 0; j-- {
			result[j] = result[j] + result[j-1]
		}
	}
	return result
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	for j := 0; j < len(triangle); j++ {
		dp[j] = triangle[n-1][j]
	}
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = dp[j] + triangle[i][j]
			if dp[j+1]+triangle[i][j] < dp[j] {
				dp[j] = dp[j+1] + triangle[i][j]
			}
		}
	}
	return dp[0]
}
