package interview_related

import "math"

func minCoins(n int) int {
	coins := []int{1, 5, 10, 50, 100, 200}

	// We only need up to n + 200
	maxAmount := n + 200

	// dp[i] stands for the min number of coins we need to pay for i money without exchange
	dp := make([]int, maxAmount+1)

	// Initialize dp with large number
	for i := 1; i <= maxAmount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[1] = 1

	for i := 2; i <= maxAmount; i++ {
		for _, coin := range coins {
			if pre := i - coin; pre >= 0 {
				dp[i] = min(dp[i], dp[pre]+1)
			}
		}
	}

	// Try all pay >= n
	ans := math.MaxInt32
	for pay := n; pay <= maxAmount; pay++ {
		ans = min(ans, dp[pay]+dp[pay-n])
	}

	return ans
}
