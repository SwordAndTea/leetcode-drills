package _511_520

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode problem No. 515

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

// leetcode problem No. 518

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	// for each coin, dp[i] represent the number of ways to form i-amount
	// by only use that coin and the coins before that coin
	// this can prevent redundant calculation
	// if we put the inner loop to the outside, there will be redundant calculation
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
