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

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visit[i][j] {
				ans++
				lands := make([][2]int, 1)
				lands[0] = [2]int{i, j}
				visit[i][j] = true
				k := 0
				for k < len(lands) {
					curIndex := lands[k]

					// left
					if curIndex[1]+1 < n && !visit[curIndex[0]][curIndex[1]+1] && grid[curIndex[0]][curIndex[1]+1] == '1' {
						lands = append(lands, [2]int{curIndex[0], curIndex[1] + 1})
						visit[curIndex[0]][curIndex[1]+1] = true
					}

					// right
					if curIndex[1]-1 >= 0 && !visit[curIndex[0]][curIndex[1]-1] && grid[curIndex[0]][curIndex[1]-1] == '1' {
						lands = append(lands, [2]int{curIndex[0], curIndex[1] - 1})
						visit[curIndex[0]][curIndex[1]-1] = true
					}

					// down
					if curIndex[0]+1 < m && !visit[curIndex[0]+1][curIndex[1]] && grid[curIndex[0]+1][curIndex[1]] == '1' {
						lands = append(lands, [2]int{curIndex[0] + 1, curIndex[1]})
						visit[curIndex[0]+1][curIndex[1]] = true
					}

					// up
					if curIndex[0]-1 >= 0 && !visit[curIndex[0]-1][curIndex[1]] && grid[curIndex[0]-1][curIndex[1]] == '1' {
						lands = append(lands, [2]int{curIndex[0] - 1, curIndex[1]})
						visit[curIndex[0]-1][curIndex[1]] = true
					}

					k++
				}
			}
		}
	}
	return ans
}
