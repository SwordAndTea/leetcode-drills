package _321_330

import "math"

// leetcode problem No. 322

func coinChange(coins []int, amount int) int {
	// dp[i] stands for the minimum number of tokens needed for i money
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if pre := i - coin; pre >= 0 && dp[pre] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[pre]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
