package _181_190

func findRepeatedDnaSequences(s string) []string {
	i, j := 0, 10
	ans := make([]string, 0)
	strMap := make(map[string]int)
	for j <= len(s) {
		str := s[i:j]
		strMap[str]++
		if strMap[str] == 2 {
			ans = append(ans, str)
		}
		i++
		j++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}

	if k >= n/2 {
		ans := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				ans += prices[i] - prices[i-1]
			}
		}
		return ans
	}

	// dp[i, j] represents the max profit up until prices[j] using at most i transactions.
	// dp[i, j] = max(dp[i, j-1], prices[j] - prices[jj] + dp[i-1, jj]), jj in range of [0, j-1]
	// dp[i][j-1] means there will no transaction at price[j]
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// For example, if j == 8, then amongst all jj == 1,2,...,7
	//The max profit could be one of the following:
	//dp[i-1][0] + prices[8] - prices[0]
	//dp[i-1][1] + prices[8] - prices[1]
	//dp[i-1][2] + prices[8] - prices[2]
	//...
	//dp[i-1][6] + prices[8] - prices[6]
	//dp[i-1][7] + prices[8] - prices[7]
	//
	//localMax is the max value amongst all:
	//dp[i-1][0] - prices[0]
	//dp[i-1][1] - prices[1]
	//dp[i-1][2] - prices[2]
	//...
	//dp[i-1][6] - prices[6]
	//dp[i-1][7] - prices[7]

	for i := 1; i <= k; i++ {
		localMax := dp[i-1][0] - prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+localMax)
			localMax = max(localMax, dp[i-1][j]-prices[j])
		}
	}
	return dp[k][n-1]
}

func rotate(nums []int, k int) {
	n := len(nums)
	if k == n {
		return
	}
	if k > n/2 {
		tmp := make([]int, n-k)
		copy(tmp, nums[0:n-k])
		copy(nums[0:k], nums[n-k:])
		copy(nums[k:], tmp)
	} else {
		tmp := make([]int, k)
		copy(tmp, nums[n-k:])
		copy(nums[k:], nums[0:n-k])
		copy(nums[0:k], tmp)
	}
}

func reverseBits(num uint32) uint32 {
	stack := make([]uint32, 32)
	i := 0
	for num != 0 {
		stack[i] = num & 1
		i++
		num >>= 1
	}
	ans := uint32(0)
	base := uint32(1)
	for i = 31; i >= 0; i-- {
		ans |= stack[i] * base
		base <<= 1
	}
	return ans
}
