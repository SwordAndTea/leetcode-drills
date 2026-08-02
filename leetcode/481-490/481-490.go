package _481_490

// leetcode problem No. 486
func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n) // dp[i][j] represent the maximum score difference
	// `current` player (player 1 or 2) can achieve for num[i:j(included)]
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			// if current player take nums[i], then it will lose the dp[i+1[j] possible points
			takeLeft := nums[i] - dp[i+1][j]

			// if current player take nums[j], then it will lose the dp[i][j-1] possible points
			takeRight := nums[j] - dp[i][j-1]
			dp[i][j] = max(takeLeft, takeRight)
		}
	}

	return dp[0][n-1] >= 0
}
