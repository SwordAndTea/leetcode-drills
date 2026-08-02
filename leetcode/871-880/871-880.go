package _871_880

// leetcode problem No. 877
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			takeLeft := piles[i] - dp[i+1][j]

			takeRight := piles[i] - dp[i][j-1]

			dp[i][j] = max(takeLeft, takeRight)
		}
	}
	return dp[0][n-1] >= 0
}
