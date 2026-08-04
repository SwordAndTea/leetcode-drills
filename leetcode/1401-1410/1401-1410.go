package _1401_1410

import "math"

// leetcode problem No. 1406
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = math.MinInt32
	}
	dp[n] = 0
	for i := n - 1; i >= 0; i-- {
		take := 0
		for j := 0; j < 3; j++ {
			takeIndex := i + j
			if takeIndex < n {
				take += stoneValue[takeIndex]
				dp[i] = max(dp[i], take-dp[takeIndex+1])
			}
		}
	}
	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	}
	return "Tie"
}
