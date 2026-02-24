package _1061_1070

import (
	"math"
	"slices"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// leetcode problem No. 1066
func assignBikes(workers [][]int, bikes [][]int) int {
	n := len(workers)
	m := len(bikes)
	// dp[i][j] represent the minimum distance to assign all j (bitmask) bikes to the first i workers
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 1<<m)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = 0 // Note, don't modify dp[i][0] and dp[0][j]
	for i := 1; i <= n; i++ {
		for j := 0; j < 1<<m; j++ {
			for k := 0; k < m; k++ { // for each bike
				if j>>k&1 == 1 { // if the k-th bike is assigned to i-th worker
					worker := workers[i-1]
					bike := bikes[k]
					d := abs(worker[0]-bike[0]) + abs(worker[1]-bike[1])
					// dp[i-1][j^(1<<k)] represent assign other chosen bikes to first i-1 workers
					dp[i][j] = min(dp[i][j], dp[i-1][j^(1<<k)]+d)
				}
			}
		}
	}

	return slices.Min(dp[n])
}
