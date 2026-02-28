package _191_200

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode problem No. 198

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	levelEnd := 1
	result := make([]int, 0)
	i := 0
	for i < levelEnd {
		result = append(result, nodes[levelEnd-1].Val)
		for i < levelEnd {
			node := nodes[i]
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
			i++
		}
		levelEnd = len(nodes)
	}
	return result
}

// leetcode problem No. 200

func numIslands(grid [][]byte) int {
	ans := 0
	m := len(grid)
	n := len(grid[0])
	var landToWater func(curI, curJ int)
	landToWater = func(curI, curJ int) {
		grid[curI][curJ] = '0'
		if curI-1 >= 0 && grid[curI-1][curJ] == '1' {
			landToWater(curI-1, curJ)
		}
		if curI+1 < m && grid[curI+1][curJ] == '1' {
			landToWater(curI+1, curJ)
		}
		if curJ-1 >= 0 && grid[curI][curJ-1] == '1' {
			landToWater(curI, curJ-1)
		}
		if curJ+1 < n && grid[curI][curJ+1] == '1' {
			landToWater(curI, curJ+1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				ans += 1
				landToWater(i, j)
			}
		}
	}
	return ans
}
