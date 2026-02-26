package _1991_2000

// leetcode problem No. 1997

func firstDayBeenInAllRooms(nextVisit []int) int {
	// dp[i] stands for the day the first visit all room from 0 to i
	// at the moment when first visit all room for 0 to i, the number of visit room i will be 1
	mod := 1_000_000_000 + 7
	dp := make([]int, len(nextVisit))
	dp[0] = 0
	dp[1] = 2
	//day := 2
	for i := 2; i < len(nextVisit); i++ {
		// (dp[i-1] - dp[nextVisit[i-1]]) stands for how many day need from visit room nextVisit[i-1] back to room i-1
		// + 2 stands for the day from room i-1 to room nextVisit[i-1] and from room i-1 to room i
		// dp[i-1] - dp[nextVisit[i-1]] may be less than 0 (because of mod), we add mod to prevent less than 0
		dp[i] = (dp[i-1] + (dp[i-1] - dp[nextVisit[i-1]] + mod) + 2) % mod
	}
	return dp[len(nextVisit)-1]
}
